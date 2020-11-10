package domain

import (
	"bus/cqrs"
	"bus/domain/commands"
	hotelevents "bus/domain/events"
	"database/sql"
	"log"
	"time"
)

type CommandsHandler struct {
	repo *cqrs.AggregateRepository
	db *sql.DB
}

func NewCommandsHandler(repo *cqrs.AggregateRepository, db *sql.DB) *CommandsHandler {
	return &CommandsHandler{repo, db}
}

func (h *CommandsHandler) HandleCommand(cmd cqrs.Command, db *sql.DB) (err error) {
	i := NewReservation(cmd.AggregateId())
	h.repo.Load(i.AggregateRoot)
	switch cmd.(type) {
	case *commands.MakeReservation:
		err = h.handleMakeReservation(i, cmd.(*commands.MakeReservation))
	case *commands.CancelReservation:
		err = h.handleCancelReservation(i, cmd.(*commands.CancelReservation))
	}
	if err != nil {
		return
	}
	uncommittedEvents := i.AggregateRoot.GetUncommittedChanges()
	err = h.repo.Save(i.AggregateRoot)

	var denormalizer = NewReservationDenormalizer(h.db)
	for _, event := range uncommittedEvents {
		denormalizer.DenormalizeEvent(event)
	}
	return
}

func (h *CommandsHandler) handleMakeReservation(i *Reservation, cmd *commands.MakeReservation) error {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	return i.MakeReservation(cmd.Id, cmd.HotelId, now, cmd.RoomType)
}

func (h *CommandsHandler) handleCancelReservation(i *Reservation, cmd *commands.CancelReservation) error {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	return i.CancelReservation(cmd.Id, cmd.HotelId, now)
}


type EventsHandler struct{}

func (h *EventsHandler) HandleEvent(e *cqrs.DomainEvent) error {
	evt := e.Event()
	switch evt.(type) {
	case *hotelevents.ReservationMade:
		return h.handleReservationMade(evt.(*hotelevents.ReservationMade))
	case *hotelevents.ReservationCanceled:
		return h.handleReservationCanceled(evt.(*hotelevents.ReservationCanceled))
	}
	return nil
}

func (h *EventsHandler) handleReservationMade(e *hotelevents.ReservationMade) error {
	log.Printf("ReservationMade: %v", e)
	return nil
}

func (h *EventsHandler) handleReservationCanceled(e *hotelevents.ReservationCanceled) error {
	log.Printf("ReservationMade: %v", e)
	return nil
}