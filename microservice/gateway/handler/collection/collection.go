package collection

import (
	"github.com/Muxi-X/forum-be/microservice/gateway/dao"
)

type Api struct {
	Dao dao.Interface
}

func New(i dao.Interface) *Api {
	api := new(Api)
	api.Dao = i
	return api
}

type CreateRequest struct {
	PostId uint32 `json:"post_id" binding:"required"`
}
