package lecturer

import (
	"net/http"
)

func (s *LecturerServer) routes() {
	// In this case s.Validate() will executed first because s.GetAllLecturer() return func not value
	s.Router.HandleFunc("/api/lecturer", s.GetAllLecturer()).Methods("GET")
	s.Router.HandleFunc("/api/lecturer/{id}", s.getLecturer()).Methods("GET")
	s.Router.HandleFunc("/api/create-lecturer", s.addLecturer()).Methods("POST")
	s.Router.Use(s.loggingMiddleware)
}

func (s *LecturerServer) Validate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// it will execute later after this line code executed because
		// naturally in Golang variable that assigned by func
		// will executed after that variable called
		w.Header().Add("Content-Type", "application/json")

		h(w, r)
	}
}

func (s *LecturerServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
