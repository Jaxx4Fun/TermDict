package rpcserver

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/base"
	"github.com/Johnny4Fun/TermDict/util/cache"
	"github.com/Johnny4Fun/TermDict/youdao"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

const (
	HTTPListenAddrAndPort = ":8808"
)

type RPCDict struct {
	lru        cache.Cache
	dict       base.Dict
	resultChan chan *base.Word
}

func DefaultRPCDict() *RPCDict {
	return NewRPCDict(DefaultLRUCache, DefaultYoudaoDict)
}

func NewRPCDict(cache cache.Cache, dict base.Dict) *RPCDict {
	return &RPCDict{
		lru:  cache,
		dict: dict,
	}
}

type Args struct {
	Word string
}

type RemoteReply struct {
	Word base.Word
}

func (s *RPCDict) LookUp(args *Args, reply *RemoteReply) error {
	wd := dictHelper(args.Word, s.lru, s.dict)
	if wd == nil {
		return fmt.Errorf("get %p", wd)
	}
	reply.Word = *wd
	return nil
}

var DefaultLRUCache = cache.NewThreadSafeLRU(cache.NewLRUCache(base.CacheCapacity))
var DefaultYoudaoDict = youdao.NewDict()

func dictHelper(word string, cache cache.Cache, dict base.Dict) *base.Word {
	// 放外面在并发的时候应该会出问题
	var cacheChan = make(chan *base.Word, 1)
	var dictChan = make(chan *base.Word, 1)
	// 异步查询
	if cache != nil {
		go func() {
			if wd := cache.Get(word); wd != nil {
				cacheChan <- wd.(*base.Word)
			}
		}()
	}

	if dict != nil {
		go func() {
			if wd := dict.LookUp(word); wd != nil {
				dictChan <- wd
			}
		}()
	}
	select {
	case wd := <-cacheChan:
		wd.From = base.MemCache
		return wd
	case wd := <-dictChan:
		wd.From = base.Online
		cache.Add(wd)
		return wd
	case <-time.After(10 * time.Second):
		log.Printf("timed out to look up %s", word)
	}

	//// 同步查询
	//if cache != nil {
	//	if wd := cache.Get(word); wd != nil {
	//		return wd.(*base.Word)
	//	}
	//}
	//
	//if dict != nil {
	//	if wd := dict.LookUp(word); wd != nil {
	//		return wd
	//	}
	//}
	//

	return nil
}

func RecoverFromCacheFiles() {
	files, err := ioutil.ReadDir(base.EnvTermDictCache)
	if err != nil {
		log.Printf("failed to read %s, %v", base.EnvTermDictRoot, err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !base.RegexCacheFileName.MatchString(file.Name()) {
			continue
		}

		buff, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Printf("failed to read %q, %v", file.Name(), err)
			continue
		}

		wd := new(base.Word)
		err = yaml.Unmarshal(buff, wd)
		if err != nil {
			log.Printf("failed to unmarshal %q, %v", file.Name(), err)
			continue
		}

		DefaultLRUCache.Add(wd)
	}
}
