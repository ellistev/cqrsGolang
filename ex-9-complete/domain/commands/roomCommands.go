package commands

import (
	"github.com/satori/go.uuid"
)

type MakeReservation struct {
	Id uuid.UUID
	RoomType string `json:"roomtype"`
	HotelId uuid.UUID `json:"hotelid"`
}

func (c *MakeReservation) AggregateId() uuid.UUID { return c.Id }

func NewReserveRoom(id uuid.UUID, hotelId uuid.UUID, roomType string) MakeReservation {
	s := MakeReservation{}
	s.Id = id
	s.HotelId = hotelId
	s.RoomType = roomType
	return s
}



type CancelReservation struct {
	Id uuid.UUID
	HotelId uuid.UUID `json:"hotelid"`
}

func NewUnReserveRoom(id uuid.UUID, hotelId uuid.UUID) CancelReservation {
	s := CancelReservation{}
	s.Id = id
	s.HotelId = hotelId
	return s
}

func (c *CancelReservation) AggregateId() uuid.UUID { return c.Id }
