package main

import _ "github.com/adel-hadadi/link-shotener/internal/common/setup"

import (
	"fmt"
	"os"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/link"
	"github.com/adel-hadadi/link-shotener/internal/common/server"
	"github.com/adel-hadadi/link-shotener/internal/link/ports"
	"github.com/adel-hadadi/link-shotener/internal/link/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	ErrConnectToDB = "error on connecting to database: %w"
	ErrLoadEnv     = "error on loading .env file: %w"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf(ErrLoadEnv, err))
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf(ErrConnectToDB, err))
	}

	app := service.NewApplication(db)

	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(app)
		link.RegisterLinkServiceServer(server, svc)
	})
}
