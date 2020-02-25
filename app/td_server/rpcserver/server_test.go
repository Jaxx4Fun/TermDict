package rpcserver_test

import (
	"github.com/Johnny4Fun/TermDict/app/td_server/rpcserver"
	"github.com/Johnny4Fun/TermDict/base"
	"net/http/httptest"
	"net/rpc"
	"strings"
	"testing"
)

func TestRPCServer_LookUp(t *testing.T) {
	word := "client"

	d := rpcserver.DefaultRPCDict()
	s := rpc.NewServer()
	s.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	err := s.Register(d)
	if err != nil {
		t.Error(err)
	}

	httpServer := httptest.NewServer(s)

	client, err := rpc.DialHTTP("tcp", strings.TrimPrefix(httpServer.URL, "http://"))
	if err != nil {
		t.Fatalf("failed to create client, %v", err)
	}

	reply := rpcserver.RemoteReply{}
	err = client.Call("RPCDict.LookUp", rpcserver.Args{Word: word}, &reply)
	if err != nil {
		t.Errorf("failed to call rpc, %v", err)
	}
	assertWordSpell(t, &reply.Word, word)
	assertWordFrom(t, &reply.Word, base.Online)

	err = client.Call("RPCDict.LookUp", rpcserver.Args{Word: word}, &reply)
	assertWordFrom(t, &reply.Word, base.MemCache)

}

func assertWordFrom(t *testing.T, word *base.Word, want base.FromStorage) {
	t.Helper()
	if from := word.From; from != want {
		t.Errorf("from %+v, expected %+v", from, want)
	}
}

func assertWordSpell(t *testing.T, word *base.Word, want string) {
	t.Helper()
	if spell := word.Spell; spell != want {
		t.Errorf("from %v, expected %v", spell, want)
	}
}
