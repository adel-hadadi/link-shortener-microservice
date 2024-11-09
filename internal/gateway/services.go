package main

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"google.golang.org/grpc"
)

type LinkService interface {
	CreateLink(ctx context.Context, originalURL string) (string, error)
	GetLink(ctx context.Context, shortURL string) (string, error)
}

type ReportService interface {
	DownloadReport(ctx context.Context, fileName string) (grpc.ServerStreamingClient[report.FileChunk], error)
}
