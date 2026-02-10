package service

import (
	"context"
	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

func (s *PostService) ListPopularTag(_ context.Context, req *pb.ListPopularTagRequest, resp *pb.Tags) error {
	logger.Info("PostService ListPopularTag")

	tags, err := s.Dao.ListPopularTags(req.Category)
	if err != nil {
		return errno.ServerErr(errno.ErrRedis, err.Error())
	}

	resp.Tags = tags

	return nil
}
