package client

import (
	"log"
	"os"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/link"
	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewReportClient() (client report.ReportServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("REPORT_GRPC_ADDR")

	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env REPORT_GRPC_ADDR")
	}

	conn, err := grpc.Dial(
		"report-grpc:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return report.NewReportServiceClient(conn), conn.Close, nil
}

func NewLinkClient() (client link.LinkServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("LINK_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env LINK_GRPC_ADDR")
	}
	log.Printf("link grpc port => %s \n", grpcAddr)

	conn, err := grpc.Dial(
		"link-grpc:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return link.NewLinkServiceClient(conn), conn.Close, nil
}
