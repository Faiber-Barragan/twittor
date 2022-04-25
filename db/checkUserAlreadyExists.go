package db

import (
	"context"
	"github.com/Faiber-Barragan/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//CheckUserAlreadyExists receive email of param and check if exists in DB
func CheckUserAlreadyExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	filter := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, filter).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
