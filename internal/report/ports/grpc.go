package ports

import (
	"context"
	"fmt"
	"io"
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

func (s GrpcServer) DownloadReport(req *report.DownloadRequest, stream report.ReportService_DownloadReportServer) error {
	buffer, err := s.app.Services.ReportService.DownloadReport(context.Background(), req.FileName)
	if err != nil {
		return err
	}

	chunkSize := 32 * 1024
	chunk := make([]byte, chunkSize)

	for {
		n, err := buffer.Read(chunk)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read buffer: %v", err)
		}
		if n == 0 {
			break
		}

		if err := stream.Send(&report.FileChunk{Content: chunk[:n]}); err != nil {
			return fmt.Errorf("failed to send chunk: %v", err)
		}
	}

	return nil
}

func protoTimestampToTime(timestamp *timestamp.Timestamp) time.Time {
	utcTime := timestamp.AsTime()

	localTime := utcTime.In(time.FixedZone("Asia/Tehran", 3*3600+30*60)) // +0330 offset

	return localTime
}
