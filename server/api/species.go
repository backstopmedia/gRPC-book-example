package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
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
