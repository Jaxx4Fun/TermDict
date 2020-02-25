// rpc server for Termdict
package main

import (
	"github.com/Johnny4Fun/TermDict/app/td_server/rpcserver"
	"github.com/Johnny4Fun/TermDict/base"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	base.Initialize()

	rpcserver.RecoverFromCacheFiles()
	s := rpcserver.DefaultRPCDict()
	rpc.Register(s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", rpcserver.HTTPListenAddrAndPort)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
