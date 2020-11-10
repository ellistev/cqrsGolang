package domain

import (
	"bus/cqrs"
	hotelevents "bus/domain/events"
	"github.com/jdextraze/go-gesclient/client"
	"github.com/satori/go.uuid"
	"time"
)

type Reservation struct {
	*cqrs.AggregateRoot
	Id                 uuid.UUID `json:"Id"`
	HotelId            uuid.UUID `json:"HotelId"`
	RoomType           string    `json:"RoomType"`
	LastUpdateDateTime time.Time `json:"LastUpdateDateTime"`
	IsCanceled         bool
}

func NewReservation(id uuid.UUID) *Reservation {
	obj := &Reservation{}
	obj.AggregateRoot = cqrs.NewAggregateRoot("Reservation", id, obj.apply)
	return obj
}

func (i *Reservation) apply(e cqrs.Event) {
	switch v := e.(type) {
	case *hotelevents.ReservationMade:
		i.onReservationMade(v)
	}
}

func (i *Reservation) MakeReservation(id uuid.UUID , hotelId uuid.UUID, lastupdateDateTime time.Time, roomType string) error {
	return i.Apply(hotelevents.NewReservationMade(id, hotelId, lastupdateDateTime, roomType))
}

func (i *Reservation) onReservationMade(e *hotelevents.ReservationMade) {
	i.Id = e.Id
}

func ( s *Reservation)DenormalizeEvent(d ReservationDenormalizer, e client.Event) {
	switch v := e.(type){
	case hotelevents.ReservationMade:
		d.OnReservationMade(v)
	}
}