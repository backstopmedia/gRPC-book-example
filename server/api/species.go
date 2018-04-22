package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSpecies(ctx context.Context, r *pb.GetSpeciesRequest) (*pb.GetSpeciesResponse, error) {
	species, err := s.db.GetSpeciesByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetSpeciesResponse{Species: species}, nil
}

func (s *Server) ListSpecies(ctx context.Context, r *pb.ListSpeciesRequest) (*pb.ListSpeciesResponse, error) {
	species, err := s.db.GetSpecies(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListSpeciesResponse{Species: species}, nil
}

func (s *Server) ValidateSpecies(ctx context.Context, r *pb.ValidateSpeciesRequest) (*pb.ValidateSpeciesResponse, error) {
	if len(r.GetName()) < 2 {
		status := status.New(codes.InvalidArgument, "invalid")
		key := &pb.InvalidKey{Key: "name", Message: "The name provided is not long enough"}

		status, err := status.WithDetails(key)
		if err != nil {
			return nil, err
		}

		return nil, status.Err()
	}

	return &pb.ValidateSpeciesResponse{}, nil
}
