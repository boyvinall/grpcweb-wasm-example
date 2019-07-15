package main

import (
	"context"
	"io"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	web "github.com/johanbrandhorst/grpcweb-wasm-example/proto"
)

func main() {
	cc, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		grpclog.Println(err)
		return
	}
	client := web.NewBackendClient(cc)
	resp, err := client.GetUser(context.Background(), &web.GetUserRequest{
		UserId: "1234",
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
	} else {
		grpclog.Println(resp)
	}
	resp, err = client.GetUser(context.Background(), &web.GetUserRequest{
		UserId: "123",
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
	} else {
		grpclog.Println(resp)
	}

	srv, err := client.GetUsers(context.Background(), &web.GetUsersRequest{
		NumUsers: 3,
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
	} else {
		for {
			user, err := srv.Recv()
			if err != nil {
				if err != io.EOF {
					st := status.Convert(err)
					grpclog.Println(st.Code(), st.Message(), st.Details())
				}
				break
			}

			grpclog.Println(user)
		}
	}

	grpclog.Println("finished")
}
