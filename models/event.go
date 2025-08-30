package models

import (
	"fmt"
	"time"

	"github.com/eduzol/event_booking_api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e Event) Save() (int64, error) {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	event, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	id, err := event.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return id, nil
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	query := `
	SELECT id, name, description, location, dateTime, user_id
	FROM events
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID); err != nil {
			fmt.Println(err)
			continue
		}
		events = append(events, e)
	}
	return events, nil
}
