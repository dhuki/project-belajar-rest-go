package lecture

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/belajarRestApi5/model"
	"github.com/belajarRestApi5/utils"
	guuid "github.com/google/uuid"
)

func (le *LectureServer) getAllLecture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (le *LectureServer) getLecture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (le *LectureServer) addLecture() http.HandlerFunc {
	type request struct { // must capitalize at first word for being captured body from request
		name       string
		Sks        int
		Topic      string
		LecturerId string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		lecture := Lecture{
			ID:          guuid.New().String(),
			Name:        req.name,
			Sks:         req.Sks,
			LectureId:   req.LecturerId,
			Topic:       req.Topic,
			IsActive:    true,
			CreatedDate: time.Now(),
		}

		result, err := lecture.AddLecture(le.DB)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		response := model.ReturnData{Success: true, Message: "Success", Data: result}
		utils.SendSuccess(w, response)
	}
}
