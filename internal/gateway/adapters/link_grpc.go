package adapters

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/link"
)

type LinkGrpc struct {
	client link.LinkServiceClient
}

func NewLinkGrpc(client link.LinkServiceClient) LinkGrpc {
	return LinkGrpc{client: client}
}

func (l LinkGrpc) CreateLink(ctx context.Context, originalURL string) (string, error) {
	linkResponse, err := l.client.Create(ctx, &link.CreateLinkRequest{OriginalLink: originalURL})
	if err != nil {
		return "", err
	}

	return linkResponse.ShortLink, nil
}

func (l LinkGrpc) GetLink(ctx context.Context, shortURL string) (string, error) {
	response, err := l.client.Get(ctx, &link.GetLinkRequest{ShortLink: shortURL})
	if err != nil {
		return "", err
	}

	return response.OriginalLink, nil
}
