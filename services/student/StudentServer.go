package student

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type StudentServer struct {
	DB     *sql.DB
	Router *mux.Router
}

func (s *StudentServer) Start() {
	s.routes()
}
