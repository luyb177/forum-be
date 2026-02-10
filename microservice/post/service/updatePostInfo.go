package service

import (
	"context"
	"strconv"

	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
	"github.com/Muxi-X/forum-be/util"
	"gorm.io/gorm"
)

func (s *PostService) UpdatePostInfo(_ context.Context, req *pb.UpdatePostInfoRequest, _ *pb.Response) error {
	logger.Info("PostService UpdatePostInfo")

	if req.Title == "" || req.Content == "" {
		return errno.ServerErr(errno.ErrBadRequest, "title and content can't be null")
	}

	post, err := s.Dao.GetPost(req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.NotFoundErr(errno.ErrItemNotFound, "post-"+strconv.Itoa(int(req.Id)))
		}

		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Domain = req.Domain
	post.CompiledContent = req.CompiledContent
	post.LastEditTime = util.GetCurrentTime()
	post.Category = req.Category
	post.Summary = req.Summary

	if err := post.Update(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
