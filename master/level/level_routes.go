package level

import (
	"github.com/gorilla/mux"
	authservice "github.com/vivaldy22/eatnfit-auth-service/proto"
	"net/http"
)

type Handler struct {
	service authservice.LevelCRUDServer
}

func NewHandler(s authservice.LevelCRUDServer, r *mux.Router) {
	handler := &Handler{s}
	r.HandleFunc("/levels", handler.GetAll).Methods(http.MethodGet)

	//prefix := r.PathPrefix("/level").Subrouter()
	//prefix.HandleFunc("", handler.Create).Methods(http.MethodPost)
	//prefix.HandleFunc("/{id}", handler.GetByID).Methods(http.MethodGet)
	//prefix.HandleFunc("/{id}", handler.Update).Methods(http.MethodPut)
	//prefix.HandleFunc("/{id}", handler.Delete).Methods(http.MethodDelete)
}