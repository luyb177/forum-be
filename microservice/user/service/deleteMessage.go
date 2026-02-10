package service

import (
	"context"

	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/user/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

// DeletePrivateMessage ... 删除用户消息
func (s *UserService) DeletePrivateMessage(_ context.Context, req *pb.DeletePrivateMessageRequest, _ *pb.Response) error {
	logger.Info("UserService DeletePrivateMessage")

	if req.Id == "" {
		if err := s.Dao.DeleteMessage(req.UserId); err != nil {
			return errno.ServerErr(errno.ErrRedis, err.Error())
		}
	} else {
		if err := s.Dao.DeleteOneMessage(req.UserId, req.Id); err != nil {
			return errno.ServerErr(errno.ErrRedis, err.Error())
		}
	}

	return nil
}
