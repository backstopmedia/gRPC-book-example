package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetStarship(context.Context, *pb.GetStarshipRequest) (*pb.GetStarshipResponse, error) {
	return nil, s.notImplementedError()
}

func (s *Server) ListStarships(context.Context, *pb.ListStarshipsRequest) (*pb.ListStarshipsResponse, error) {
	return nil, s.notImplementedError()
}
