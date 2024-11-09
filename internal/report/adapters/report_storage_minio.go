package adapters

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
)

type ReportStorageMinio struct {
	client *minio.Client
}

func NewReportStorageMinio(client *minio.Client) ReportStorageMinio {
	return ReportStorageMinio{
		client: client,
	}
}

func (m ReportStorageMinio) GetObject(ctx context.Context, filePath string) (*bytes.Buffer, error) {
	object, err := m.client.GetObject(ctx, os.Getenv("MINIO_REPORT_BUCKET"), filePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, object); err != nil {
		return nil, err
	}

	return &buf, nil
}

func (m ReportStorageMinio) PutObject(ctx context.Context, content *bytes.Buffer, filePath string) error {
	_, err := m.client.PutObject(ctx, os.Getenv("MINIO_REPORT_BUCKET"), filePath, content, int64(content.Len()), minio.PutObjectOptions{
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	})

	return err
}
