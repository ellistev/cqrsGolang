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



type CancelReservation struct {
	Id uuid.UUID
	HotelId uuid.UUID `json:"hotelid"`
}



func (c *CancelReservation) AggregateId() uuid.UUID { return c.Id }
