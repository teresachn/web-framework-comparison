package repository

import (
	"context"
	"os"
	config "tchen/web-framework-comparison/config"
	entity "tchen/web-framework-comparison/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	MONGO_MOVIE_INFO_DB = os.Getenv("MONGO_MOVIE_INFO_DB")
	tableName           = "movie_infos"
)

func AddMovieInfo(ctx context.Context, config config.Configuration, movieInfo *entity.MovieInfo) error {
	insertInfo, err := config.Client.Database(MONGO_MOVIE_INFO_DB).Collection(tableName).InsertOne(ctx, movieInfo)
	if err != nil {
		return err
	}

	movieInfo.ID = insertInfo.InsertedID.(primitive.ObjectID)
	return nil
}

func GetAllMovieInfo(ctx context.Context, config config.Configuration) ([]entity.MovieInfo, error) {
	cursor, err := config.Client.Database(MONGO_MOVIE_INFO_DB).Collection(tableName).Find(ctx, bson.D{})
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
