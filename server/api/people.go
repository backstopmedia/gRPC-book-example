package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetPerson(ctx context.Context, r *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	person, err := s.db.GetPersonByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPersonResponse{Person: person}, nil
}

func (s *Server) ListPeople(ctx context.Context, r *pb.ListPeopleRequest) (*pb.ListPeopleResponse, error) {
	people, err := s.db.GetPeople(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListPeopleResponse{People: people}, nil
}
