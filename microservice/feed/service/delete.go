package service

import (
	"context"
	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/feed/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

func (s *FeedService) Delete(_ context.Context, req *pb.Request, _ *pb.Response) error {
	logger.Info("FeedService Delete")

	err := s.Dao.Delete(req.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
