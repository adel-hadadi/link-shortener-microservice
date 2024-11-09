package adapters

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"google.golang.org/grpc"
)

type ReportGrpc struct {
	client report.ReportServiceClient
}

func NewReportGrpc(client report.ReportServiceClient) ReportGrpc {
	return ReportGrpc{client: client}
}

func (l ReportGrpc) DownloadReport(ctx context.Context, fileName string) (grpc.ServerStreamingClient[report.FileChunk], error) {
	stream, err := l.client.DownloadReport(ctx, &report.DownloadRequest{
		FileName: fileName,
	})

	return stream, err
}
