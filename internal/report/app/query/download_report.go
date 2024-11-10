package query

import (
	"bytes"
	"context"
)

type DownloadReportHandler struct {
	storage reportStorage
}

type reportStorage interface {
	GetObject(ctx context.Context, filePath string) (*bytes.Buffer, error)
}

func NewDownloadReportHnadler(storage reportStorage) DownloadReportHandler {
	return DownloadReportHandler{storage: storage}
}

func (h DownloadReportHandler) Handle(ctx context.Context, filePath string) (*bytes.Buffer, error) {
	content, err := h.storage.GetObject(ctx, filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}
