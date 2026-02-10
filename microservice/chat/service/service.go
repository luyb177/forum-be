package service

import (
	"github.com/Muxi-X/forum-be/microservice/chat/dao"
)

// ChatService ... 聊天服务
type ChatService struct {
	Dao dao.Interface
}

func New(i dao.Interface) *ChatService {
	service := new(ChatService)
	service.Dao = i
	return service
}
