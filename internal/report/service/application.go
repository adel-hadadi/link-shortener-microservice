package service

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/report/adapters"
	"github.com/adel-hadadi/link-shotener/internal/report/app"
	"github.com/adel-hadadi/link-shotener/internal/report/app/service"
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

	repo := adapters.NewMongoClientRepository(mongoClient)

	return app.Application{
			Services: app.Services{
				ReportService: service.NewReportService(repo),
			},
		}, func() {
			_ = mongoClient.Disconnect(context.Background())
		}
}
