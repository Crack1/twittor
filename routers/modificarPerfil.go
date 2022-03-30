package routers

import (
	"encoding/json"
	"github.com/Crack1/twittor/db"
	"github.com/Crack1/twittor/models"
	"net/http"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}
	var status bool
	status, err = db.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrion un error al intentar modificar el Registro"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro de usuario"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
