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
//virtual-workshop ex-6 implement a switch statement for getting the *commands.CancelReservation command here
	}
	return "", nil
}

func (f *Factory) CreateEvent(name string) (Event, error) {
	switch name {
	case "ReservationMade":
		return &hotelevents.ReservationMade{}, nil
	}
	return nil, nil
}

func (f *Factory) GetEventType(evt Event) (string, error) {
	switch evt.(type) {
	case hotelevents.ReservationMade:
		return "ReservationMade", nil
	}
	return "", nil
}