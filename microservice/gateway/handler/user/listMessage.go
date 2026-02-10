package user

import (
	"context"

	"github.com/Muxi-X/forum-be/client"
	"github.com/Muxi-X/forum-be/log"
	. "github.com/Muxi-X/forum-be/microservice/gateway/handler"
	"github.com/Muxi-X/forum-be/microservice/gateway/util"
	pb "github.com/Muxi-X/forum-be/microservice/user/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListMessage ... 获取 user message list
// @Summary list user message api
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} ListMessageResponse
// @Router /user/message/list [get]
func ListMessage(c *gin.Context) {
	log.Info("User ListMessage function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	userId := c.MustGet("userId").(uint32)

	listReq := &pb.ListMessageRequest{
		UserId: userId,
	}

	listResp, err := client.UserClient.ListMessage(context.TODO(), listReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendMicroServiceResponse(c, nil, listResp, ListMessageResponse{})
}

// ListPrivateMessage ... 获取 user private_message list
// @Summary list user private_message api
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} ListMessageResponse
// @Router /user/private_message/list [get]
func ListPrivateMessage(c *gin.Context) {
	log.Info("User ListPrivateMessage function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	userId := c.MustGet("userId").(uint32)

	listReq := &pb.ListMessageRequest{
		UserId: userId,
	}

	listResp, err := client.UserClient.ListPrivateMessage(context.TODO(), listReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendMicroServiceResponse(c, nil, listResp, ListMessageResponse{})
}
