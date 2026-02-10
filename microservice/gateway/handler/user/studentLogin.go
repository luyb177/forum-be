package user

import (
	"context"

	"github.com/Muxi-X/forum-be/client"
	"github.com/Muxi-X/forum-be/log"
	. "github.com/Muxi-X/forum-be/microservice/gateway/handler"
	"github.com/Muxi-X/forum-be/microservice/gateway/util"
	pb "github.com/Muxi-X/forum-be/microservice/user/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// StudentLogin ... 学生登录
// @Summary 学生登录 api
// @Description login the student-forum
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body StudentLoginRequest true "login_request"
// @Success 200 {object} StudentLoginResponse
// @Router /auth/login/student [post]
func StudentLogin(c *gin.Context) {
	log.Info("User StudentLogin function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req StudentLoginRequest
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 login
	loginReq := &pb.StudentLoginRequest{
		StudentId: req.StudentId,
		Password:  req.Password,
	}

	loginResp, err := client.UserClient.StudentLogin(context.TODO(), loginReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendMicroServiceResponse(c, nil, loginResp, StudentLoginResponse{})
}
