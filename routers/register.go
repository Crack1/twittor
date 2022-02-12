package routers

import (
	"encoding/json"
	"github.com/Crack1/twittor/db"
	"github.com/Crack1/twittor/models"
	"net/http"
)

/* Register is the function that create in the DB the user registration */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El Password tiene que tener mas de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := db.CheckIfUserAlreadyExist(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe ese email", 400)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
