// td的rpc客户端
package main

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/app/td_server/rpcserver"
	"log"
	"net/rpc"
	"os"
)

var (
//word  = flag.String("word", "", "word to loop up")
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	word := os.Args[1]
	//directLookup()
	rpcLookup(word)
}

func rpcLookup(word string) {
	client, err := rpc.DialHTTP("tcp", rpcserver.HTTPListenAddrAndPort)
	if err != nil {
		log.Fatalf("failed to dial rpc @%s, %v", rpcserver.HTTPListenAddrAndPort, err)
	}
	args := rpcserver.Args{"you"}
	reply := rpcserver.RemoteReply{}
	err = client.Call("RPCDict.LookUp", &args, &reply)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply.Word)
}
