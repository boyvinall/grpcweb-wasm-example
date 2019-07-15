// +build js

package main

import (
	"io/ioutil"
	"syscall/js"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var document js.Value

type DivWriter js.Value

func (d DivWriter) Write(p []byte) (n int, err error) {
	node := document.Call("createElement", "div")
	node.Set("innerHTML", string(p))
	js.Value(d).Call("appendChild", node)
	return len(p), nil
}

func init() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "target")
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(DivWriter(div), ioutil.Discard, ioutil.Discard))
}

const endpoint = ""

var opts = []grpc.DialOption{}
