package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Faiber-Barragan/twittor/middlew"
	"github.com/Faiber-Barragan/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handlers I set my port, the handler and I put to listing my server
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
