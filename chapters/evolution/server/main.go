package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"errors"
	"log"
	"net"

	pb "github.com/backstopmedia/gRPC-book-example/chapters/evolution/server/proto"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, new(MyService))

	log.Print("Starting RPC server on port 8080...")
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to setup tcp listener: %v", err)
	}

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type MyService struct{}

func (s *MyService) UpdateThing(ctx context.Context, r *pb.UpdateThingRequest) (*pb.UpdateThingResponse, error) {
	if r.Mask != nil {
		// Update only fields in r.Mask.Paths
		log.Print(r.Mask.Paths)
	}

	// thing, err := findThingByID(r.Thing.Id)
	// assuming this was returned
	err := errors.New("Thing not found")
	if err != nil {
		statusErr := status.New(codes.NotFound, err.Error())
		statusErr, _ = statusErr.WithDetails(&pb.ThingNotFoundError{
			RequestedId: r.Thing.Id,
			Reason:      "Thing not found, or you do not have access to it.",
		})

		return nil, statusErr.Err()
	}

	return &pb.UpdateThingResponse{Thing: r.Thing}, nil
}
