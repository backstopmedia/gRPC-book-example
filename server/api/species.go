package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetSpecies(context.Context, *pb.GetSpeciesRequest) (*pb.GetSpeciesResponse, error) {
	return nil, s.notImplementedError()
}

func (s *Server) ListSpecies(context.Context, *pb.ListSpeciesRequest) (*pb.ListSpeciesResponse, error) {
	return nil, s.notImplementedError()
}
