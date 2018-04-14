package api

import (
	"time"

	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
)

const (
	millenniumFalconID = "rJe-kah5NsM"
)

func (s *Server) ListStarshipActions(req *google_protobuf1.Empty, stream pb.Starwars_ListStarshipActionsServer) error {
	starship, err := s.db.GetStarshipByID(stream.Context(), millenniumFalconID)
	if err != nil {
		return err
	}

	actions := []pb.StarshipAction_Action{
		pb.StarshipAction_TOOKOFF,
		pb.StarshipAction_HYPERDRIVE,
		pb.StarshipAction_LANDED,
		pb.StarshipAction_HIDING_IN_A_MOUTH,
	}

	for _, action := range actions {
		response := &pb.StarshipAction{
			Action:   action,
			Starship: starship,
		}

		if err := stream.Send(response); err != nil {
			return err
		}

		time.Sleep(time.Second * 3)
	}

	return nil
}
