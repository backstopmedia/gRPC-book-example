package main

import (
	"log"

	pb "github.com/backstopmedia/gRPC-book-example/chapters/streaming/server/proto"
)

// DatabaseService is an implementation of the Database service in database.proto
type DatabaseService struct{}

// Search returns a stream of matching search results
func (db *DatabaseService) Search(r *pb.SearchRequest, s pb.Database_SearchServer) error {
	responses := []string{
		"Highest ranked content",
		"Some ranked content",
		"Some ranked content",
		"Lowest ranked content",
	}

	for idx, resp := range responses {
		result := &pb.SearchResponse{MatchedTerm: r.Term, Rank: int32(idx + 1), Content: resp}

		if err := s.Send(result); err != nil {
			log.Printf("Error sending message to the client: %v", err)
			return err
		}
	}

	return nil
}
