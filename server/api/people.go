package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetPerson(context.Context, *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	return nil, s.notImplementedError()
}

func (s *Server) ListPeople(context.Context, *pb.ListPeopleRequest) (*pb.ListPeopleResponse, error) {
	return nil, s.notImplementedError()
}
