package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Crack1/twittor/middleware"
	"github.com/Crack1/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Drivers, port is setted */
func Drivers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.CheckDB(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middleware.CheckDB(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
