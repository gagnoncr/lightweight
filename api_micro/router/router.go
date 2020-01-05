package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"api_micro/middelware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "nothing to see here")
	})
	router.HandleFunc("/api/health-check", middleware.LogHandler(healthcheck)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/set", middleware.CreateDeployment).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deployments", middleware.GetAllDeployments).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/deleteDeployment/{id}", middleware.DeleteDeployment).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllDeployments", middleware.DeleteAllDeployments).Methods("DELETE", "OPTIONS")

	return router
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "healthy")
	w.WriteHeader(http.StatusAccepted)
}