package lecturer

import (
	"database/sql"
	"os"

	"github.com/gorilla/mux"
)

type LecturerServer struct {
	DB     *sql.DB
	Router *mux.Router
	Log    *os.File
}

func (s *LecturerServer) Start() {
	s.routes()
}
