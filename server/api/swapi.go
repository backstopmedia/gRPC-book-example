package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represents the server implementation of the SW API.
type Server struct {
}

func (s *Server) notImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
	return st.Err()
}
