package lecture

import (
	"net/http"
)

func (le *LectureServer) routes() {
	le.Router.HandleFunc("/api/lecture", le.getAllLecture()).Methods("GET")
	le.Router.HandleFunc("/api/lecture/{id}", le.getLecture()).Methods("GET")
	le.Router.HandleFunc("/api/create-lecture", le.addLecture()).Methods("POST")
	le.Router.Use(le.loggingMiddleware)
}

func (le *LectureServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
