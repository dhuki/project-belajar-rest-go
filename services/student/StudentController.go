package student

import (
	"encoding/json"
	"net/http"
	"time"

	guuid "github.com/google/uuid"

	"github.com/belajarRestApi5/model"
	"github.com/belajarRestApi5/utils"
)

func (s *StudentServer) getAllStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student Student

		result, err := student.getAllStudent(s.DB)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		var response model.ReturnData

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

func (s *StudentServer) getStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *StudentServer) addStudent() http.HandlerFunc {

	type request struct { // must capitalize at first word for being captured body from request
		FirstName string
		LastName  string
		Dob       string
		Pob       string
	}
	return func(w http.ResponseWriter, r *http.Request) {

		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		student := Student{
			ID:          guuid.New().String(),
			FirstName:   req.FirstName,
			LastName:    req.LastName,
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

		student.DOB = dob

		result, err := student.addStudent(s.DB)
		if err != nil {
			response := model.ReturnData{Success: false, Message: err.Error()}
			utils.SendError(w, http.StatusInternalServerError, response)
			return
		}

		response := model.ReturnData{Success: true, Message: "Success", Data: result}
		utils.SendSuccess(w, response)
	}
}
