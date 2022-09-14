package entity

import (
	"time"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieInfo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"movie_name"`
	ReleaseDate time.Time          `bson:"release_date"`
}
