package student

import "net/http"

func (s *StudentServer) routes() {
	// In this case s.Validate() will executed first because s.GetAllLecturer() return func not value
	s.Router.HandleFunc("/api/students", s.getAllStudent()).Methods("GET")
	s.Router.HandleFunc("/api/students/{id}", s.getStudent()).Methods("GET")
	s.Router.HandleFunc("/api/students", s.addStudent()).Methods("POST")
	s.Router.Use(s.loggingMiddleware)
}

func (s *StudentServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
