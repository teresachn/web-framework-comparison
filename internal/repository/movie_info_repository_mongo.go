package repository

import (
	"context"
	"os"
	config "tchen/web-framework-comparison/config"
	entity "tchen/web-framework-comparison/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MONGO_MOVIE_INFO_DB = os.Getenv("MONGO_MOVIE_INFO_DB")
)

type MovieInfoRepositoryMongo struct {
	database  *mongo.Database
	tableName string
}

func NewMovieInfoRepositoryMongo(config config.MongoConfig) MovieInfoRepository {
	return &MovieInfoRepositoryMongo{
		database:  config.Database,
		tableName: MONGO_MOVIE_INFO_DB,
	}
}

func (m MovieInfoRepositoryMongo) AddMovieInfo(ctx context.Context, movieInfo *entity.MovieInfo) error {
	insertInfo, err := m.database.Collection(m.tableName).InsertOne(ctx, movieInfo)
	if err != nil {
		return err
	}

	movieInfo.ID = insertInfo.InsertedID.(primitive.ObjectID)
	return nil
}

func (m MovieInfoRepositoryMongo) GetAllMovieInfo(ctx context.Context) ([]entity.MovieInfo, error) {
	cursor, err := m.database.Collection(m.tableName).Find(ctx, bson.D{})
	if err != nil {
		return []entity.MovieInfo{}, err
	}

	var result []entity.MovieInfo
	err = cursor.All(ctx, &result)
	if err != nil {
		return []entity.MovieInfo{}, err
	}

	return result, nil
}
