package db

import (
	"context"
	"github.com/Crack1/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

/* InsertRegister  */
func InsertRegister(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("golang-api")
	col := db.Collection("usuarios")

	u.Password, _ := PasswordEncrypt(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil{
		return "",false,  err
	}

	ObjID, _ := result.InsertedID(primitive.ObjectID)
	return ObjID.String(), true, nil
}
