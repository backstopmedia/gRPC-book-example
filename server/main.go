package main

import (
	"google.golang.org/grpc"

	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/backstopmedia/gRPC-book-example/server/db"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

const (
	defaultDataFile = "data.json"
	defaultPort     = 8080
)

var (
	dataFile string
	port     int
)

func main() {
	flag.IntVar(&port, "port", defaultPort, "The port to listen on")
	flag.StringVar(&dataFile, "file", defaultDataFile, "The json file to use as a database")
	flag.Parse()

	log.Printf("Starting RPC server on port %d...", port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to setup tcp listener: %v", err)
	}

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalf("unable to load data file: %v", err)
	}

	db, err := db.NewJSONProvider(data)
	if err != nil {
		log.Fatalf("unable to parse data file: %v", err)
	}

	interceptorOpt := grpc.UnaryInterceptor(api.Interceptors())
	s := grpc.NewServer(interceptorOpt)
	pb.RegisterStarfriendsServer(s, api.New(db))

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
