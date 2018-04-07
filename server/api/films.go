package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetFilm(context.Context, *pb.GetFilmRequest) (*pb.GetFilmResponse, error) {
	return nil, s.notImplementedError()
}

func (s *Server) ListFilms(context.Context, *pb.ListFilmsRequest) (*pb.ListFilmsResponse, error) {
	return nil, s.notImplementedError()
}
