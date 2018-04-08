package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/backstopmedia/gRPC-book-example/server/db"
)

// Server represents the server implementation of the SW API.
type Server struct {
	db db.Provider
}

func New(provider db.Provider) *Server {
	return &Server{db: provider}
}

func (s *Server) notImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
	return st.Err()
}
