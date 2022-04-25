package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//User is a model of user of the DB of MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	BirthDay  time.Time          `bson:"birth_day,omitempty" json:"birth_day,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebSite   string             `bson:"web_site" json:"web_site,omitempty"`
}
