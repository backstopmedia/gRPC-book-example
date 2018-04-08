package api

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

func (s *Server) GetFilm(ctx context.Context, r *pb.GetFilmRequest) (*pb.GetFilmResponse, error) {
	film, err := s.db.GetFilmByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetFilmResponse{Film: film}, nil
}

func (s *Server) ListFilms(ctx context.Context, r *pb.ListFilmsRequest) (*pb.ListFilmsResponse, error) {
	films, err := s.db.GetFilms(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListFilmsResponse{Films: films}, nil
}
