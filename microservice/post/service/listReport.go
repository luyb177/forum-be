package service

import (
	"context"

	logger "github.com/Muxi-X/forum-be/log"
	pb "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/pkg/errno"
)

func (s *PostService) ListReport(_ context.Context, req *pb.ListReportRequest, resp *pb.ListReportResponse) error {
	logger.Info("PostService ListReport")

	reports, err := s.Dao.ListReport(req.Offset, req.Limit, req.LastId, req.Pagination)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	resp.Reports = reports

	return nil
}
