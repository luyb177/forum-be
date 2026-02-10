package service

import (
	"context"

	logger "github.com/Muxi-X/forum-be/log"
	"github.com/Muxi-X/forum-be/microservice/chat/dao"
	pb "github.com/Muxi-X/forum-be/microservice/chat/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

// Create 发送消息
func (s *ChatService) Create(_ context.Context, req *pb.CreateRequest, _ *pb.Response) error {
	logger.Info("CharService Create")

	data := &dao.ChatData{
		Content:  req.Content,
		Time:     req.Time,
		Receiver: req.TargetUserId,
		Sender:   req.UserId,
		TypeName: req.TypeName,
	}

	err := s.Dao.Create(data)

	if err != nil {
		return errno.ServerErr(errno.ErrRedis, err.Error())
	}

	return nil
}
