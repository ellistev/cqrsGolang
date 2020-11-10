package requests

import (
	"github.com/satori/go.uuid"
)

type IRequest interface{
	CommandId() uuid.UUID
}

type BaseIRequest struct {
	Id uuid.UUID `json:"id"`
}

func (c BaseIRequest) CommandId() uuid.UUID{
	return c.Id
}

type ReserveRoomRequest struct{
	BaseIRequest
	RoomType string `json:"roomtype"`
	HotelId uuid.UUID `json:"hotelid"`
}


type UnReserveRoomRequest struct{
	BaseIRequest
	HotelId uuid.UUID `json:"hotelid"`
}
