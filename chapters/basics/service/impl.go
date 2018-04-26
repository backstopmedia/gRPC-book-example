package service

import (
  "time"

  "github.com/golang/protobuf/ptypes"
  "github.com/golang/protobuf/ptypes/timestamp"
  "golang.org/x/net/context"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  "../proto"
)

// To start with, we'll hardcode the database of films.
var films = []*proto.Film{
  &proto.Film{
    Id:          "4",
    Title:       "A New Hope",
    Director:    "George Lucas",
    Producer:    "Gary Kurtz, Rick McCallum",
    ReleaseDate: toProto(1977, 5, 25),
  },
  &proto.Film{
    Id:          "5",
    Title:       "The Empire Strikes Back",
    Director:    "Irvin Kershner",
    Producer:    "Gary Kurtz, Rick McCallum",
    ReleaseDate: toProto(1980, 5, 17),
  },
  &proto.Film{
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
  return &proto.GetFilmResponse{Film: film}, nil
}

// ListFilms returns a list of all known films.
func (s *StarfriendsImpl) ListFilms(ctx context.Context,
    req *proto.ListFilmsRequest) (*proto.ListFilmsResponse, error) {

  return &proto.ListFilmsResponse{Films: films}, nil
}

// compile-type check that our new type provides the
// correct server interface
var _ proto.StarfriendsServer = (*StarfriendsImpl)(nil)

