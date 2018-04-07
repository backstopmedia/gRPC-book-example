package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetVehicle(context.Context, *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	return nil, s.notImplementedError()
}

func (s *Server) ListVehicles(context.Context, *pb.ListVehiclesRequest) (*pb.ListVehiclesResponse, error) {
	return nil, s.notImplementedError()
}
