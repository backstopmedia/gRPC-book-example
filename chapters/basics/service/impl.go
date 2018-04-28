package service

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"../proto"
)

// To start with, we'll hardcode the database of films.
var films = []*proto.Film{
	{
		Id:          "4",
		Title:       "A New Hope",
		Director:    "George Lucas",
		Producer:    "Gary Kurtz, Rick McCallum",
		ReleaseDate: toProto(1977, 5, 25),
	},
	{
		Id:          "5",
		Title:       "The Empire Strikes Back",
		Director:    "Irvin Kershner",
		Producer:    "Gary Kurtz, Rick McCallum",
		ReleaseDate: toProto(1980, 5, 17),
	},
	{
		Id:          "6",
		Title:       "Return of the Jedi",
		Director:    "Richard Marquand",
		Producer:    "Howard G. Kazanjian, George Lucas, Rick McCallum",
		ReleaseDate: toProto(1983, 5, 25),
	},
}

func toProto(year, month, day int) *timestamp.Timestamp {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		panic(err)
	}
	return ts
}

type StarfriendsImpl struct {
}

// GetFilm queries a film by ID or returns an error if not found.
func (s *StarfriendsImpl) GetFilm(ctx context.Context,
	req *proto.GetFilmRequest) (*proto.GetFilmResponse, error) {

	respHdrs := metadata.New(map[string]string{
		"Who":     "starfriends-server",
		"Version": "v1",
	})
	grpc.SendHeader(ctx, respHdrs)

	start := time.Now()
	defer func() {
		respTrlrs := metadata.Pairs("duration", time.Since(start).String())
		grpc.SetTrailer(ctx, respTrlrs)
	}()

	if reqHdrs, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("Request headers: %v", reqHdrs)
		// Must use all-lower-case keys to query metadata
		if who, ok := reqHdrs["who"]; ok {
			// who is a slice of strings; just use the first
			log.Printf("Received request from %s", who[0])
		}
	}

	var film *proto.Film
	for _, f := range films {
		if f.Id == req.Id {
			film = f
			break
		}
	}
	if film == nil {
		return nil, status.Errorf(codes.NotFound, "no film with id %q", req.Id)
	}
	resp := &proto.GetFilmResponse{Film: film}

	return resp, nil
}

// ListFilms returns a list of all known films.
func (s *StarfriendsImpl) ListFilms(ctx context.Context,
	req *proto.ListFilmsRequest) (*proto.ListFilmsResponse, error) {

	return &proto.ListFilmsResponse{Films: films}, nil
}

// compile-type check that our new type provides the
// correct server interface
var _ proto.StarfriendsServer = (*StarfriendsImpl)(nil)
