package comment

import (
	"context"
	"strconv"

	"github.com/Muxi-X/forum-be/client"
	"github.com/Muxi-X/forum-be/log"
	. "github.com/Muxi-X/forum-be/microservice/gateway/handler"
	"github.com/Muxi-X/forum-be/microservice/gateway/util"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/model"
	"github.com/Muxi-X/forum-be/pkg/constvar"
	"github.com/Muxi-X/forum-be/pkg/errno"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Get ... 获取评论
// @Summary 获取评论 api
// @Tags comment
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param comment_id path int true "comment_id"
// @Success 200 {object} Comment
// @Router /comment/{comment_id} [get]
func (a *Api) Get(c *gin.Context) {
	log.Info("Comment Get function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	userId := c.MustGet("userId").(uint32)

	ok, err := model.Enforce(userId, constvar.Comment, id, constvar.Read)
	if err != nil {
		SendError(c, errno.ErrCasbin, nil, err.Error(), GetLine())
		return
	}

	if !ok {
		SendError(c, errno.ErrPermissionDenied, nil, "权限不足", GetLine())
		return
	}

	getReq := &pb.Request{
		UserId: userId,
		Id:     uint32(id),
	}

	getResp, err := client.PostClient.GetComment(context.TODO(), getReq)
	if err != nil {
		SendError(c, err, getResp, "", GetLine())
		return
	}

	SendMicroServiceResponse(c, nil, getResp, Comment{})
}
