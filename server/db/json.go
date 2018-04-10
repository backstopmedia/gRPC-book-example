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

func (p *jsonProvider) GetSpeciesByID(ctx context.Context, id string) (*pb.Species, error) {
	for _, s := range p.db.Species {
		if s.Id == id {
			return s, nil
		}
	}

	return nil, fmt.Errorf("species with id '%s' not found", id)
}

func (p *jsonProvider) GetSpecies(ctx context.Context) ([]*pb.Species, error) {
	return p.db.Species, nil
}

func (p *jsonProvider) GetStarshipByID(ctx context.Context, id string) (*pb.Starship, error) {
	for _, s := range p.db.Starships {
		if s.Id == id {
			return s, nil
		}
	}

	return nil, fmt.Errorf("starship with id '%s' not found", id)
}

func (p *jsonProvider) GetStarships(ctx context.Context) ([]*pb.Starship, error) {
	return p.db.Starships, nil
}

func (p *jsonProvider) GetVehicleByID(ctx context.Context, id string) (*pb.Vehicle, error) {
	for _, v := range p.db.Vehicles {
		if v.Id == id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("vehicle with id '%s' not found", id)
}

func (p *jsonProvider) GetVehicles(ctx context.Context) ([]*pb.Vehicle, error) {
	return p.db.Vehicles, nil
}
