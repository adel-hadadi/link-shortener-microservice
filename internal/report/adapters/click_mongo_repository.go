package adapters

import (
	"context"
	"github.com/adel-hadadi/link-shotener/internal/report/app/service"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClickRepository struct {
	mongoClient *mongo.Client
}

type Click struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	ShortURL  string    `bson:"short_url" json:"short_url"`
	ClickedAt time.Time `bson:"clicked_at" json:"clicked_at"`
}

type ClickCount struct {
	ShortURL string `json:"short_url" bson:"_id"`
	Count    int    `json:"count" bson:"clickCount"`
}

func NewMongoClientRepository(client *mongo.Client) MongoClickRepository {
	return MongoClickRepository{
		mongoClient: client,
	}
}

func (r MongoClickRepository) Create(ctx context.Context, shortURL string, clickedAt time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	collection := r.mongoClient.Database("report").Collection("clicks")
	_, err := collection.InsertOne(ctx, Click{
		ShortURL: shortURL,

		ClickedAt: clickedAt,
	})

	return err
}

func (r MongoClickRepository) GetLastHourClicks(ctx context.Context) ([]service.ClickCount, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	oneHourAgo := time.Now().UTC().Add(-1 * time.Hour)

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"clicked_at", bson.D{
				{"$gte", oneHourAgo},
				{"$lte", time.Now()},
			}},
		}}},
		{{"$group", bson.D{
			{"_id", "$short_url"},
			{"clickCount", bson.D{
				{"$sum", 1},
			}},
		}}},
	}

	collection := r.mongoClient.Database("report").Collection("clicks")

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []ClickCount
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return r.clickCountModelToService(result), nil
}

func (r MongoClickRepository) clickCountModelToService(counts []ClickCount) []service.ClickCount {
	res := make([]service.ClickCount, 0, len(counts))

	for c := range counts {
		res = append(res, service.ClickCount{
			ShortURL: counts[c].ShortURL,
			Count:    counts[c].Count,
		})
	}

	return res
}

func (r MongoClickRepository) clickModelsToService(clicks []Click) []service.Click {
	res := make([]service.Click, 0, len(clicks))
	for click := range clicks {
		res = append(res, service.Click{
			ShortURL:  clicks[click].ShortURL,
			ClickedAt: clicks[click].ClickedAt,
		})
	}

	return res
}
