package ports

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/link"
	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/adel-hadadi/link-shotener/internal/link/app/service"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func (s GrpcServer) Create(ctx context.Context, req *link.CreateLinkRequest) (*link.LinkResponse, error) {
	shortLink, err := s.app.Services.LinkService.Create(ctx, service.CreateLink{OriginalLink: req.OriginalLink})
	if err != nil {
		return nil, err
	}

	return &link.LinkResponse{
		ShortLink:    shortLink,
		OriginalLink: req.OriginalLink,
	}, nil
}

func (s GrpcServer) Get(ctx context.Context, req *link.GetLinkRequest) (*link.LinkResponse, error) {
	l, err := s.app.Services.LinkService.Get(ctx, req.ShortLink)
	if err != nil {
		return nil, err
	}

	return &link.LinkResponse{
		OriginalLink: l.OriginalLink,
	}, nil
}
