package lecture

import (
	"database/sql"
	"time"

	"github.com/belajarRestApi5/services/lecturer"
)

type Lecture struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Sks         int    `json:"sks"`
	LectureId   string `json:"lectureId"`
	Lecturer    *lecturer.Lecturer
	Topic       string    `json:"topic"`
	CreatedDate time.Time `json:"createdDate"`
	IsActive    bool      `json:"isActive"`
	Ordinal     int32     `json:"ordinal"`
}

func (le *Lecture) AddLecture(db *sql.DB) (Lecture, error) {
	rows, err := db.Query("Insert into lecture (id, name, sks, lecturer_id, topic, is_active, created_date) values($1, $2, $3, $4, $5, $6, $7)",
		le.ID, le.Name, le.Sks, le.LectureId, le.Topic, le.IsActive, le.CreatedDate)

	if err != nil {
		return Lecture{}, err
	}

	defer rows.Close() //will execute after all line code executed

	for rows.Next() {
		err := rows.Scan(&le.ID, &le.Name, &le.Sks, &le.LectureId, &le.Topic, &le.Ordinal, &le.IsActive, &le.CreatedDate)
		if err != nil {
			return Lecture{}, err
		}
	}

	return *le, nil
}
