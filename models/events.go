package models

import (
	"go-event-api/db"
	"log"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.Id = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return []Event{}, err
	}
	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ? 
	WHERE id = ?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)

	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Id)

	return err
}

func (e *Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?,?)`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Id, userId)

	if err != nil {
		return err
	}
	return nil
}

func (e *Event) DeleteRegistration(userId int64) error {
	query := `
	DELETE FROM registrations WHERE event_id = ? AND user_id = ?	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Id, userId)

	if err != nil {
		return err
	}
	return nil
}
