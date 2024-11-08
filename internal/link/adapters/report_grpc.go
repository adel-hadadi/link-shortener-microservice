package adapters

import (
	"context"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReportGrpc struct {
	client report.ReportServiceClient
}

func NewReportGrpc(client report.ReportServiceClient) ReportGrpc {
	return ReportGrpc{client: client}
}

func (s ReportGrpc) LinkClicked(ctx context.Context, shortURL string, clickedAt time.Time) error {
	_, err := s.client.LinkClicked(
		ctx,
		&report.LinkClickedRequest{
			ShortUrl:  shortURL,
			ClickedAt: timestamppb.New(clickedAt.UTC()),
		},
	)

	return err
}
