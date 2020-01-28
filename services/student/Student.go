package student

import (
	"database/sql"
	"time"
)

type Student struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"fristName"`
	LastName    string    `json:"last_name"`
	DOB         time.Time `json:"date_of_birth"`
	POB         string    `json:"place_of_birthday"`
	Ordinal     int32     `json:"ordinal"`
	IsActive    bool      `json:"isActive"`
	CreatedDate time.Time `json:"createdDate"`
}

func (s *Student) getAllStudent(db *sql.DB) ([]Student, error) {

	rows, err := db.Query("Select * from student")
	defer rows.Close() //will execute after all line code executed
	if err != nil {
		return nil, err
	}

	var students []Student

	//if rows.Next() { // check if query has next value

	for rows.Next() {

		if students == nil {
			students = []Student{} //define a slice -> FYI slice more lightweight than array because its using pointer
		}

		err := rows.Scan(&s.ID, &s.FirstName, &s.LastName,
			&s.DOB, &s.POB, &s.Ordinal, &s.IsActive, &s.CreatedDate) //Scan copies the columns from the matched row (in database) into the values pointed (in struct) at by its destination
		if err != nil {
			return nil, err
		}

		students = append(students, *s)
	}

	return students, nil
}

func (s *Student) addStudent(db *sql.DB) (Student, error) {

	rows, err := db.Query("Insert into student (id, first_name, last_name, dob, pob, is_active, created_date) values($1, $2, $3, $4, $5, $6, $7)",
		s.ID, s.FirstName, s.LastName, s.DOB, s.POB, s.IsActive, s.CreatedDate)
	if err != nil {
		return Student{}, err
	}

	defer rows.Close() //will execute after all line code executed

	for rows.Next() {
		err := rows.Scan(&s.ID, &s.FirstName, &s.LastName, &s.DOB, &s.POB, &s.Ordinal, &s.IsActive, &s.CreatedDate)
		if err != nil {
			return Student{}, err
		}
	}

	return *s, nil
}
