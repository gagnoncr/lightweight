package router

import (
	"fmt"
	"net/http"
	"web_micro/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/", http.FileServer(http.Dir("./static")).ServeHTTP).Methods("GET", "OPTIONS")
	router.HandleFunc("/health-check", middleware.LogHandler(healthCheck)).Methods("GET", "OPTIONS")

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "healthy")
	w.WriteHeader(http.StatusAccepted)
}