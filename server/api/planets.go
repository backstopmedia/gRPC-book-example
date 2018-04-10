package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetPlanet(context.Context, *pb.GetPlanetRequest) (*pb.GetPlanetResponse, error) {
	return nil, s.notImplementedError()
}

func (s *Server) ListPlanets(context.Context, *pb.ListPlanetsRequest) (*pb.ListPlanetsResponse, error) {
	return nil, s.notImplementedError()
}
