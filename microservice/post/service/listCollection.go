package service

import (
	"context"

	logger "github.com/Muxi-X/forum-be/log"
	"github.com/Muxi-X/forum-be/microservice/post/dao"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/pkg/constvar"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

func (s *PostService) ListCollection(_ context.Context, req *pb.ListPostPartInfoRequest, resp *pb.ListPostPartInfoResponse) error {
	logger.Info("PostService ListCollections")

	postIds, err := s.Dao.ListCollectionByUserId(req.TargetUserId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	var filter = &dao.PostModel{}

	domain, err := s.GetUserDomain(req.UserId)
	if err != nil {
		return errno.ServerErr(errno.ErrRPC, err.Error())
	}
	if domain == constvar.NormalDomain {
		filter.Domain = domain
	}

	posts, err := s.Dao.ListPostInfoByPostIds(postIds, filter, req.Offset, req.Limit, req.LastId, req.Pagination)
	if err != nil {
		return errno.ServerErr(errno.ErrListPostInfoByPostIds, err.Error())
	}

	resp.Posts = posts

	return nil
}
