package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/4ddx/apiB/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	log.Println("Listening on: ", s.addr)
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
