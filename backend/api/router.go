package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi-zen/sea-turtle/backend/api/middlewares"
	"github.com/yoshi-zen/sea-turtle/backend/controllers"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	c := controllers.NewProblemController(db)

	r.HandleFunc("/problem/list", c.GetProblemListHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/detail", c.GetProblemDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/add", c.PostProblemHandler).Methods(http.MethodPost)

	r.Use(middlewares.Logging)

	return r
}
