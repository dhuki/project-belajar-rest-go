package config

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/belajarRestApi5/services/lecturer"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

func (s *Server) Initialize() {

	// err := godotenv.Load() //question why we should import this while not used but it make works
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// username := os.Getenv("db.username")
	// password := os.Getenv("db.password")
	// dbName := os.Getenv("db.name")
	// dbHost := os.Getenv("db.host")

	// username := "ijjfgiyjcbrznd"
	// password := "028ccc705a477aabf05483a69471b3cd349122d598495748014d446c7aad41dd"
	// dbName := "d4ppbeeimehina"
	// dbHost := "ec2-174-129-255-21.compute-1.amazonaws.com"

	// //(format without printing)
	// dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	// fmt.Println(dbURI)

	dbURI := "postgres://ijjfgiyjcbrznd:028ccc705a477aabf05483a69471b3cd349122d598495748014d446c7aad41dd@ec2-174-129-255-21.compute-1.amazonaws.com:5432/d4ppbeeimehina"

	conn, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	s.DB = conn
	s.Router = mux.NewRouter()
}

func (s *Server) SetupServer() {

	serverLecturer := &lecturer.LecturerServer{
		DB:     s.DB,
		Router: s.Router,
	}

	serverLecturer.Start()
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}

func MainServerStart() {
	mainServer := Server{}

	mainServer.Initialize()
	mainServer.SetupServer()
	mainServer.Run()
}
