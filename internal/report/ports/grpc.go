package ports

import (
	"context"
	"log"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"github.com/adel-hadadi/link-shotener/internal/report/app"
	"github.com/adel-hadadi/link-shotener/internal/report/app/service"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func (s GrpcServer) LinkClicked(ctx context.Context, req *report.LinkClickedRequest) (*empty.Empty, error) {
	clickedAt := protoTimestampToTime(req.ClickedAt)

	log.Println("clicked at => ", clickedAt)

	err := s.app.Services.ReportService.LinkClicked(
		ctx,
		service.LinkClickedReq{
			Link:      req.ShortUrl,
			ClickedAt: clickedAt,
		},
	)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func protoTimestampToTime(timestamp *timestamp.Timestamp) time.Time {
	utcTime := timestamp.AsTime()

	localTime := utcTime.In(time.FixedZone("Asia/Tehran", 3*3600+30*60)) // +0330 offset

	return localTime
}
