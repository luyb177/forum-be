package post

import (
	"context"

	"github.com/Muxi-X/forum-be/client"
	"github.com/Muxi-X/forum-be/log"
	. "github.com/Muxi-X/forum-be/microservice/gateway/handler"
	"github.com/Muxi-X/forum-be/microservice/gateway/util"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListPopularTag ... 获取热门tags
// @Summary list 热门tags api
// @Description 降序
// @Tags post
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param category query string false "category"
// @Success 200 {object} []string
// @Router /post/popular_tag [get]
func (a *Api) ListPopularTag(c *gin.Context) {
	log.Info("Post ListPopularTag function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	category := c.DefaultQuery("category", "")

	req := &pb.ListPopularTagRequest{
		Category: category,
	}

	resp, err := client.PostClient.ListPopularTag(context.TODO(), req)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, resp.Tags)
}
