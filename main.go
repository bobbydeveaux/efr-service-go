package main

import (
	"github.com/bobbydeveaux/efr-service-go/grpc"
	"github.com/bobbydeveaux/efr-service-go/rest"
)

func main() {
	go grpc.Serve()
	rest.Serve()

}
