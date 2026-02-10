package service

import (
	"context"
	logger "github.com/Muxi-X/forum-be/log"
	"github.com/Muxi-X/forum-be/microservice/post/dao"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/model"
	"github.com/Muxi-X/forum-be/pkg/constvar"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

func (s *PostService) DeleteCollection(_ context.Context, req *pb.Request, _ *pb.Response) error {
	logger.Info("PostService DeleteCollection")

	collection := &dao.CollectionModel{
		Id: req.Id,
	}

	if err := s.Dao.DeleteCollection(collection); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if err := model.DeletePermission(req.UserId, constvar.Collection, req.Id, constvar.Write); err != nil {
		return errno.ServerErr(errno.ErrCasbin, err.Error())
	}

	go func() {
		if err := s.Dao.ChangePostScore(req.Id, -constvar.CollectionScore); err != nil {
			logger.Error(errno.ErrChangeScore.Error(), logger.String(err.Error()))
		}
	}()

	return nil
}
