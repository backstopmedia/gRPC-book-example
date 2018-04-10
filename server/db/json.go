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

func (p *jsonProvider) GetPersonByID(ctx context.Context, id string) (*pb.Person, error) {
	for _, person := range p.db.People {
		if person.Id == id {
			return person, nil
		}
	}

	return nil, fmt.Errorf("person with id '%s' not found", id)
}

func (p *jsonProvider) GetPeople(ctx context.Context) ([]*pb.Person, error) {
	return p.db.People, nil
}

func (p *jsonProvider) GetPlanetByID(ctx context.Context, id string) (*pb.Planet, error) {
	for _, planet := range p.db.Planets {
		if planet.Id == id {
			return planet, nil
		}
	}

	return nil, fmt.Errorf("planet with id '%s' not found", id)
}

func (p *jsonProvider) GetPlanets(ctx context.Context) ([]*pb.Planet, error) {
	return p.db.Planets, nil
}
