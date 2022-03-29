package routers

import (
	"errors"
	"github.com/Crack1/twittor/db"
	"github.com/Crack1/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterDeDesarrollo_Evides")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := db.CheckIfUserAlreadyExist(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Invalido")
	}
	return claims, false, string(""), err
}
