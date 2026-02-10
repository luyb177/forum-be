package service

import "github.com/Muxi-X/forum-be/microservice/user/dao"

// UserService ... 用户服务
type UserService struct {
	Dao dao.Interface
}

func New(i dao.Interface) *UserService {
	service := new(UserService)
	service.Dao = i
	return service
}
