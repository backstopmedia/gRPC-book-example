package api

import (
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

const (
	millenniumFalconID = "rJe-kah5NsM"
)

func (s *Server) ListStarshipActions(req *pb.ListStarshipActionsRequest, stream pb.Starfriends_ListStarshipActionsServer) error {
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
	}

	return nil
}
