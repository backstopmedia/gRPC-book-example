package db

import (
	"github.com/golang/protobuf/jsonpb"

	"bytes"
	"context"
	"fmt"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type jsonProvider struct {
	db *pb.DB
}

func NewJSONProvider(data []byte) (Provider, error) {
	var db pb.DB
	if err := jsonpb.Unmarshal(bytes.NewReader(data), &db); err != nil {
		return nil, err
	}

	return &jsonProvider{db: &db}, nil
}

func (p *jsonProvider) GetFilmByID(ctx context.Context, id string) (*pb.Film, error) {
	for _, f := range p.db.Films {
		if f.Id == id {
			return f, nil
		}
	}

	return nil, fmt.Errorf("film with id '%s' not found", id)
}

func (p *jsonProvider) GetFilms(ctx context.Context) ([]*pb.Film, error) {
	return p.db.Films, nil
}
