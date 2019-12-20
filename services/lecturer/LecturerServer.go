package lecturer

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type LecturerServer struct {
	DB     *sql.DB
	Router *mux.Router
}

func (s *LecturerServer) Start() {
	s.routes()
}
