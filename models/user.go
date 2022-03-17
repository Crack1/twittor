package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usuario /* User Model */
type Usuario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Lastname  string             `bson:"lastName" json:"lastname,omitempty"`
	Birthday  time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biografia string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	WebSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
