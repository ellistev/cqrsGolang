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


//virtual workshop ex-5
//implement the structs to represent the CancelReservation command