package adapters

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"google.golang.org/grpc"
)

type ReportGrpc struct {
	client report.ReportServiceClient
}

func NewReportGrpc(client report.ReportServiceClient) ReportGrpc {
	return ReportGrpc{client: client}
}

func (g ReportGrpc) DownloadReport(ctx context.Context, fileName string) (grpc.ServerStreamingClient[report.FileChunk], error) {
	stream, err := g.client.DownloadReport(ctx, &report.DownloadRequest{
		FileName: fileName,
	})

	return stream, err
}

func (g ReportGrpc) LinkClicked(ctx context.Context, shortURL string, clickedAt time.Time) error {
	_, err := g.client.LinkClicked(
		ctx,
		&report.LinkClickedRequest{
			ShortUrl:  shortURL,
			ClickedAt: timestamppb.New(clickedAt.UTC()),
		},
	)

	return err
}
