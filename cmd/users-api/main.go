package main

import (
	"context"
	"net/http"

	gw "github.com/avasapollo/users-api/web"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	endpoint = "localhost:9090"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	mux.Handle("POST", "/v1/users/")
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", mux)
}
