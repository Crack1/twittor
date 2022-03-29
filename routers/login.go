package routers

import (
	"encoding/json"
	"fmt"
	"github.com/Crack1/twittor/db"
	"github.com/Crack1/twittor/jwt"
	"github.com/Crack1/twittor/models"
	"net/http"
	"time"
)

// Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Password invalidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email de usuario es requerido", 400)
		return
	}

	documento, existe := db.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Password invalidos", 400)
		return
	}

	fmt.Println(documento)
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token Correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
