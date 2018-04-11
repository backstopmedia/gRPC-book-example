package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetPlanet(ctx context.Context, r *pb.GetPlanetRequest) (*pb.GetPlanetResponse, error) {
	planet, err := s.db.GetPlanetByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPlanetResponse{Planet: planet}, nil
}

func (s *Server) ListPlanets(ctx context.Context, r *pb.ListPlanetsRequest) (*pb.ListPlanetsResponse, error) {
	planet, err := s.db.GetPlanets(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListPlanetsResponse{Planets: planet}, nil
}
