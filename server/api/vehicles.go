package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetVehicle(ctx context.Context, r *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	vehicle, err := s.db.GetVehicleByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetVehicleResponse{Vehicle: vehicle}, nil
}

func (s *Server) ListVehicles(ctx context.Context, r *pb.ListVehiclesRequest) (*pb.ListVehiclesResponse, error) {
	vehicles, err := s.db.GetVehicles(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListVehiclesResponse{Vehicles: vehicles}, nil
}
