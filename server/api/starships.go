package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetStarship(ctx context.Context, r *pb.GetStarshipRequest) (*pb.GetStarshipResponse, error) {
	starship, err := s.db.GetStarshipByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetStarshipResponse{Starship: starship}, nil
}

func (s *Server) ListStarships(ctx context.Context, r *pb.ListStarshipsRequest) (*pb.ListStarshipsResponse, error) {
	starships, err := s.db.GetStarships(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListStarshipsResponse{Starships: starships}, nil
}
