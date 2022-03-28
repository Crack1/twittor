package jwt

import (
	"github.com/Crack1/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MasterDeDesarrollo_Evides")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Name,
		"apellido":         t.Lastname,
		"fecha_nacimiento": t.Birthday,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.WebSite,
		"__id":             t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
