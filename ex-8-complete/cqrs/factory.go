package cqrs

import (
	"bus/domain/commands"
	hotelevents "bus/domain/events"
	"errors"
)

type Factory struct{}

func (f *Factory) CreateCommand(name string) (Command, error) {
	return nil, errors.New("Not implemented")
}

func (f *Factory) GetCommandType(cmd Command) (string, error) {
	switch cmd.(type) {
	case *commands.MakeReservation:
		return "MakeReservation", nil
	case *commands.CancelReservation:
		return "CancelReservation", nil
	}
	return "", nil
}

func (f *Factory) CreateEvent(name string) (Event, error) {
	switch name {
	case "ReservationMade":
		return &hotelevents.ReservationMade{}, nil
	case "ReservationCanceled":
		return &hotelevents.ReservationCanceled{}, nil
	}
	return nil, nil
}

func (f *Factory) GetEventType(evt Event) (string, error) {
	switch evt.(type) {
	case hotelevents.ReservationMade:
		return "ReservationMade", nil
	case hotelevents.ReservationCanceled:
		return "ReservationCanceled", nil
	}
	return "", nil
}