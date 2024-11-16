package domain

import (
	"errors"
	"time"
)

type Status int

const (
	Active Status = iota + 1
	Cancelled
	Completed
)

type Reservation struct {
	id        string
	userID    string
	evseUID   string
	startTime time.Time
	endTime   time.Time
	status    Status
}

func NewReservation(id, userID, evseUID string, startTime, endTime time.Time, status Status) (*Reservation, error) {
	if id == "" || userID == "" || evseUID == "" {
		return nil, errors.New("id, userID, and evseUID cannot bi empty")
	}
	if startTime.After(endTime) {
		return nil, errors.New("startTime must be before endTime")
	}
	if status < Active || status > Completed {
		return nil, errors.New("invalid status")
	}

	return &Reservation{
		id:        id,
		userID:    userID,
		evseUID:   evseUID,
		startTime: startTime,
		endTime:   endTime,
		status:    status,
	}, nil
}
