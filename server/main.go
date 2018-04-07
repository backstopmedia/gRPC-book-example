package main

import (
	"google.golang.org/grpc"

	"flag"
	"fmt"
	"log"
	"net"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

const defaultPort = 8080

var port int

func main() {
	flag.IntVar(&port, "port", defaultPort, "The port to listen on")
	flag.Parse()

	log.Printf("Starting RPC server on port %d...", port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to setup tcp listener: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStarwarsServer(s, new(api.Server))

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
