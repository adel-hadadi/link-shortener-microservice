package ports

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/link"
	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/adel-hadadi/link-shotener/internal/link/app/command"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func (s GrpcServer) Create(ctx context.Context, req *link.CreateLinkRequest) (*link.LinkResponse, error) {
	err := s.app.Commands.GenerateShortURL.Handle(ctx, command.GenerateShortURL{
		OriginalURL: req.OriginalLink,
	})
	if err != nil {
		return nil, err
	}

	shortURL, err := s.app.Queries.RetrieveShortURL.Handle(ctx, req.OriginalLink)
	if err != nil {
		return nil, err
	}

	return &link.LinkResponse{
		ShortLink:    shortURL,
		OriginalLink: req.OriginalLink,
	}, nil
}

func (s GrpcServer) Get(ctx context.Context, req *link.GetLinkRequest) (*link.LinkResponse, error) {
	originalURL, err := s.app.Queries.RetrieveOriginalURL.Handle(ctx, req.ShortLink)
	if err != nil {
		return nil, err
	}

	return &link.LinkResponse{
		OriginalLink: originalURL,
	}, nil
}
