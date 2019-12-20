package lecturer

import (
	"database/sql"
	"time"
)

type Lecturer struct {

	// this process is marshalling (the process of generating a JSON string from a data)
	// or it can be use to unmarshalling (the process of generating a the act of parsing JSON to a data structure)
	ID          string     `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DOB         *time.Time `json:"date_of_birth"` // it's make that variable can assign null
	POB         string     `json:"place_of_birth"`
	Ordinal     int32      `json:"ordinal"`
	IsActive    bool       `json:"is_active"`
	CreatedDate time.Time  `json:"created_date"`
}

func (l *Lecturer) getAllLecturer(db *sql.DB) ([]Lecturer, error) {
	// First, and most important, does the method need to modify the receiver struct? If it does, the receiver must be a pointer.

	rows, err := db.Query("Select * from lecturer")
	defer rows.Close() //will execute after all line code executed
	if err != nil {
		return nil, err
	}

	var lecturers []Lecturer

	//if rows.Next() { // check if query has next value

	for rows.Next() {

		if lecturers == nil {
			lecturers = []Lecturer{} //define a slice -> FYI slice more lightweight than array because its using pointer
		}

		err := rows.Scan(&l.ID, &l.FirstName, &l.LastName,
			&l.DOB, &l.POB, &l.Ordinal, &l.IsActive, &l.CreatedDate) //Scan copies the columns from the matched row (in database) into the values pointed (in struct) at by its destination
		if err != nil {
			return nil, err
		}

		lecturers = append(lecturers, *l)
	}

	return lecturers, nil

	//}
}

func (l *Lecturer) getLecturer(db *sql.DB) (Lecturer, error) {
	// First, and most important, does the method need to modify the receiver's struct? If it does, the receiver must be a pointer.

	//Receiver parameters can be passed as either values or pointers of the base type (l.ID == params["id"])
	rows, err := db.Query("Select * from lecturer where id = $1", l.ID)
	if err != nil {
		return Lecturer{}, err
	}

	defer rows.Close() //will execute after all line code executed

	for rows.Next() {
		err := rows.Scan(&l.ID, &l.FirstName, &l.LastName, &l.DOB, &l.POB, &l.Ordinal, &l.IsActive, &l.CreatedDate)
		if err != nil {
			return Lecturer{}, err
		}
	}

	return *l, nil

}

func (l *Lecturer) AddLecturer(db *sql.DB) (Lecturer, error) {

	rows, err := db.Query("Insert into lecturer (id, first_name, last_name, dob, pob, is_active, created_date) values($1, $2, $3, $4, $5, $6, $7)", l.ID, l.FirstName, l.LastName, l.DOB, l.POB, l.IsActive, l.CreatedDate)

	if err != nil {
		return Lecturer{}, err
	}

	defer rows.Close() //will execute after all line code executed

	for rows.Next() {
		err := rows.Scan(&l.ID, &l.FirstName, &l.LastName, &l.DOB, &l.POB, &l.Ordinal, &l.IsActive, &l.CreatedDate)
		if err != nil {
			return Lecturer{}, err
		}
	}

	return *l, nil
}
