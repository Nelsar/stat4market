package db

import (
	"context"
	"log"

	"stat4market.com/models"
)

func UniqueEventType() *[]models.Event {
	connection, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	rows, err := connection.Query(ctx, "select eventType from events GROUP BY eventType HAVING COUNT(eventType) > 1000")

	if err != nil {
		log.Fatalln(err)
	}

	events := []models.Event{}

	for rows.Next() {
		event := models.Event{}

		err := rows.Scan(&event.EventType)
		if err != nil {
			log.Fatalln(err)
		}

		events = append(events, event)
	}

	return &events
}

func AddEvent(event models.Event) error {
	connection, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	err = connection.Exec(ctx, "INSERT INTO events VALUES(?, ?, ?, ?, ?)", event.EventId, event.EventType, event.UserId, event.EventTime)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUserEvents() *[]models.Event {
	connection, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	rows, err := connection.Query(ctx, "select userID from events group by userID having count(eventType) = 3")

	events := []models.Event{}

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		event := models.Event{}

		err := rows.Scan(&event.UserId)
		if err != nil {
			log.Fatalln(err)
		}

		events = append(events, event)
	}

	return &events
}

func GetStartOfMonthEvents() *[]models.Event {
	connection, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	rows, err := connection.Query(ctx, "select eventType from events where eventTime = toStartOfMonth(now())")

	events := []models.Event{}

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		event := models.Event{}

		err := rows.Scan(&event.UserId)
		if err != nil {
			log.Fatalln(err)
		}

		events = append(events, event)
	}

	return &events
}

func GetEvents(event models.Event) *[]models.Event {
	connection, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	rows, err := connection.Query(ctx, "select eventID, eventType, userID, eventTime, payload from events where eventType = ? and eventTime = ?", &event.EventType, &event.EventTime)

	events := []models.Event{}

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		event := models.Event{}

		err := rows.Scan(&event.UserId)
		if err != nil {
			log.Fatalln(err)
		}

		events = append(events, event)
	}

	return &events
}
