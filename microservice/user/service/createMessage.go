package service

import (
	"context"
	"encoding/json"
	"strconv"

	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/user/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"

	"github.com/google/uuid"
)

// CreateMessage ... 创建用户消息
func (s *UserService) CreateMessage(_ context.Context, req *pb.CreateMessageRequest, _ *pb.Response) error {
	logger.Info("UserService CreateMessage")

	if err := s.Dao.CreateMessage(0, req.Message); err != nil {
		return errno.ServerErr(errno.ErrRedis, err.Error())
	}

	return nil
}

func (s *UserService) CreatePrivateMessage(_ context.Context, req *pb.CreatePrivateMessageRequest, _ *pb.Response) error {
	logger.Info("UserService CreatePrivateMessage")

	uid := uuid.New().String()
	message, _ := json.Marshal(map[string]string{
		"id":          uid,
		"send_userid": strconv.Itoa(int(req.SendId)),
		"content":     req.Content,
		"post_id":     strconv.Itoa(int(req.PostId)),
		"comment_id":  strconv.Itoa(int(req.CommentId)),
		"type":        req.Type,
	})

	if err := s.Dao.CreateMessage(req.ReceiveId, string(message)); err != nil {
		return errno.ServerErr(errno.ErrRedis, err.Error())
	}

	return nil
}
