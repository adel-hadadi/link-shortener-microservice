package app

import (
	"github.com/adel-hadadi/link-shotener/internal/link/app/command"
	"github.com/adel-hadadi/link-shotener/internal/link/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	GenerateShortURL command.GenerateShortURLHandler
}

type Queries struct {
	RetrieveOriginalURL query.RetrieveOriginalURLHandler
	RetrieveShortURL    query.RetrieveShortURLHandler
}
