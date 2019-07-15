// +build !js

package main

import (
	"crypto/tls"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const endpoint = "localhost:10000"

var opts = []grpc.DialOption{
	grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true, // nolint: gosec
	})),
}

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
}
