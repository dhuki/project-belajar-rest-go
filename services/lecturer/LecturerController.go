package lecturer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/belajarRestApi5/model"
	"github.com/belajarRestApi5/utils"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *LecturerServer) GetAllLecturer() http.HandlerFunc {
	// First, and most important, does the method need to modify the receiver struct? If it does, the receiver must be a pointer.
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.ReturnData

		var lecturer Lecturer
		result, error := lecturer.getAllLecturer(s.DB)

		if error != nil {
			response.Success = false
			response.Message = error.Error()
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		if result == nil {
			response.Success = false
			response.Message = "Data is empty"
		} else {
			response.Success = true
			response.Message = "Success"
			response.Data = result

		}
		utils.SendSuccess(w, response)
	}
}

func (s *LecturerServer) getLecturer() http.HandlerFunc {
	// First, and most important, does the method need to modify the receiver struct? If it does, the receiver must be a pointer.
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var response model.ReturnData

		var lecturer Lecturer

		//lecturer.ID, _ = strconv.Atoi(params["id"]) //convert from string to int (strconv.Atoi return two params -> value, error)

		//lecturer.ID = "8c62e9d6-6744-4a89-9ab0-23fe7fd55498"

		value, error := strconv.Atoi(params["id"])
		if error != nil {
			log.Println("Lecturer Controller - getLecturer", error.Error())
			response.Success = false
			response.Message = error.Error()
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		fmt.Println(value)

		result, error := lecturer.getLecturer(s.DB)

		if error != nil {
			response.Success = false
			response.Message = error.Error()
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		response.Success = true
		response.Message = "Success"
		response.Data = result

		utils.SendSuccess(w, response)
	}
}

func (s *LecturerServer) addLecturer() http.HandlerFunc {
	// First, and most important, does the method need to modify the receiver struct? If it does, the receiver must be a pointer.

	type request struct { // must capitalize at first word for being captured body from request
		FirstName string
		LastName  string
		Dob       string
		Pob       string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.ReturnData

		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		lecturer := Lecturer{
			ID:        guuid.New().String(),
			FirstName: req.FirstName,
			LastName:  req.LastName,
			// DOB: func() *time.Time { // it's called annonymous func
			// 	t, err := time.Parse("2006-01-02", req.Dob)
			// 	if err != nil {
			// 		response := model.ReturnData{Success: false, Message: err.Error()}
			// 		utils.SendError(w, http.StatusInternalServerError, response)
			// 	} else {
			// 		return &t
			// 	}
			// 	return nil
			// }(),
			POB:         req.Pob,
			IsActive:    true,
			CreatedDate: time.Now(),
		}

		dob, err := utils.ConvertStringtoDate(utils.FORMAT_DATE, req.Dob)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		lecturer.DOB = &dob

		result, err := lecturer.AddLecturer(s.DB)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		response.Success = true
		response.Message = "Success"
		response.Data = result

		utils.SendSuccess(w, response)
	}
}
