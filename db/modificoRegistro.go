package db

import (
	"context"
	"github.com/Crack1/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("golang-api")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}
	if len(u.Lastname) > 0 {
		registro["lastname"] = u.Name
	}
	registro["birthday"] = u.Birthday
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.WebSite) > 0 {
		registro["website"] = u.WebSite
	}

	updateString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}