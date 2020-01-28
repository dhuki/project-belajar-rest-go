package lecture

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type LectureServer struct {
	DB     *sql.DB
	Router *mux.Router
}

func (le *LectureServer) Start() {
	le.routes()
}
