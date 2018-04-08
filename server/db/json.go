package db

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type jsonDB struct {
	Films []*film
}

type film struct {
	ID           string
	Title        string
	EpisodeID    int32
	OpeningCrawl string
	Director     string
	Producer     string
	ReleaseDate  string
}

type jsonProvider struct {
	db *jsonDB
}

func NewJSONProvider(data []byte) (Provider, error) {
	var db jsonDB
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	return &jsonProvider{db: &db}, nil
}

func (p *jsonProvider) GetFilmByID(ctx context.Context, id string) (*pb.Film, error) {
	for _, f := range p.db.Films {
		if f.ID == id {
			return toProtoFilm(f), nil
		}
	}

	return nil, fmt.Errorf("film with id '%s' not found", id)
}

func (p *jsonProvider) GetFilms(ctx context.Context) ([]*pb.Film, error) {
	films := make([]*pb.Film, len(p.db.Films))
	for i, f := range p.db.Films {
		films[i] = toProtoFilm(f)
	}

	return films, nil
}

func toProtoFilm(f *film) *pb.Film {
	return &pb.Film{
		EpisodeId:    f.EpisodeID,
		Director:     f.Director,
		Id:           f.ID,
		OpeningCrawl: f.OpeningCrawl,
		Producer:     f.Producer,
		ReleaseDate:  f.ReleaseDate,
		Title:        f.Title,
	}
}
