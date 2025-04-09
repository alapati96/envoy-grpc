package main

import (
	"context"
	"fmt"
	"net"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	auth_pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AuthServer struct{}

func (server *AuthServer) Check(ctx context.Context, request *auth_pb.CheckRequest) (*auth_pb.CheckResponse, error) {
	// block if path is /private
	// path := request.Attributes.Request.Http.Path[1:]
	// fmt.Println("Path: ", path)

	// if path == "private" {
	// 	fmt.Println("blocked private request")
	// 	return nil, fmt.Errorf("private request not allowed")
	// }

	// allow all other requests
	// add custom headers
	headers := map[string]string{
		"x-custom-header": "hello-word",
	}

	fmt.Println("Appending hello-world")
	return &auth_pb.CheckResponse{
		HttpResponse: &auth_pb.CheckResponse_OkResponse{
			OkResponse: &auth_pb.OkHttpResponse{
				Headers: SetHeaders(headers),
			},
		},
	}, nil
}

func SetHeaders(headers map[string]string) []*corev3.HeaderValueOption {
	var headerValueOptions []*corev3.HeaderValueOption
	for key, value := range headers {
		headerValueOptions = append(headerValueOptions, &corev3.HeaderValueOption{
			Header: &corev3.HeaderValue{
				Key:   key,
				Value: value,
			},
		})
	}

	return headerValueOptions
}

func main() {
	// struct with check method
	// endPoint := fmt.Sprintf("localhost:%d", 9000)
	listen, _ := net.Listen("tcp", ":9000")

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	// register envoy proto server
	server := &AuthServer{}
	auth_pb.RegisterAuthorizationServer(grpcServer, server)

	fmt.Println("Server started at port 9000")
	grpcServer.Serve(listen)
}
