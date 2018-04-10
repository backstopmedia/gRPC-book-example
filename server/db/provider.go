package db

import (
	"context"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type Provider interface {
	GetFilmByID(ctx context.Context, id string) (*pb.Film, error)
	GetFilms(ctx context.Context) ([]*pb.Film, error)

	GetPersonByID(ctx context.Context, id string) (*pb.Person, error)
	GetPeople(ctx context.Context) ([]*pb.Person, error)

	GetPlanetByID(ctx context.Context, id string) (*pb.Planet, error)
	GetPlanets(ctx context.Context) ([]*pb.Planet, error)

	GetSpeciesByID(ctx context.Context, id string) (*pb.Species, error)
	GetSpecies(ctx context.Context) ([]*pb.Species, error)
}
