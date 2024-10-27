package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adel-hadadi/link-shotener/internal/common/server"
	"github.com/adel-hadadi/link-shotener/internal/link/ports"
	"github.com/adel-hadadi/link-shotener/internal/link/service"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

const ErrConnectToDB = "error on connecting to database: %w"

func main() {
	godotenv.Load()

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

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(app),
			router,
		)
	})
}
