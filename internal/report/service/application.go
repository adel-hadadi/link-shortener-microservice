package service

import (
	"context"
	"fmt"
	"os"

	"github.com/adel-hadadi/link-shotener/internal/report/adapters"
	"github.com/adel-hadadi/link-shotener/internal/report/app"
	"github.com/adel-hadadi/link-shotener/internal/report/app/service"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewApplication() (app.Application, func()) {
	mongoOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	mongoOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	mongoClient, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		panic(err)
	}

	minioClient, err := minio.New(os.Getenv("MINIO_URL"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET"), ""),
		Secure: false,
	})
	if err != nil {
		panic(fmt.Errorf("error on connecting to minio: %w", err))
	}

	storage := adapters.NewReportStorageMinio(minioClient)

	repo := adapters.NewMongoClientRepository(mongoClient)

	return app.Application{
			Services: app.Services{
				ReportService: service.NewReportService(repo, storage),
			},
		}, func() {
			_ = mongoClient.Disconnect(context.Background())
		}
}
