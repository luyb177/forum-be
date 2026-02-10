package service

import (
	"context"

	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/model"
	"github.com/Muxi-X/forum-be/pkg/constvar"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

func (s *PostService) DeleteItem(_ context.Context, req *pb.DeleteItemRequest, _ *pb.Response) error {
	logger.Info("PostService DeleteItem")

	var err error

	if req.TypeName == constvar.Post {
		err = s.Dao.DeletePost(req.Id)
	} else if req.TypeName == constvar.Comment {
		err = s.Dao.DeleteComment(req.Id)
	} else {
		return errno.ServerErr(errno.ErrBadRequest, "wrong TypeName")
	}
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if err := model.DeletePermission(req.UserId, req.TypeName, req.Id, constvar.Write); err != nil {
		return errno.ServerErr(errno.ErrCasbin, err.Error())
	}

	return nil
}
