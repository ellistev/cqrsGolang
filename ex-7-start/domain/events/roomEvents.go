package hotelevents

import (
	"github.com/satori/go.uuid"
	"time"
)

type ReservationMade struct{
	Id uuid.UUID `json:"Id"`
	HotelId uuid.UUID `json:"HotelId"`
	RoomType string `json:"RoomType"`
	LastUpdateDateTime time.Time `json:"LastUpdateDateTime"`
}

func NewReservationMade(id uuid.UUID, hotelId uuid.UUID, lastUpdateTime time.Time, roomType string) ReservationMade {
	s := ReservationMade{}
	s.Id = id
	s.HotelId = hotelId
	s.RoomType = roomType
	s.LastUpdateDateTime = lastUpdateTime
	return s
}

type ReservationCanceled struct{
	Id uuid.UUID `json:"Id"`
	HotelId uuid.UUID `json:"HotelId"`
	ReservationId uuid.UUID `json:"ReservationId"`
	LastUpdateDateTime time.Time `json:"LastUpdateDateTime"`
}

//virtual-workshop ex-7 your IDE should indicate that this method is now in use, one the CancelReservation method in the room.go aggregate properly
// results in a new ReservationCanceled event
func NewReservationCanceled(id uuid.UUID, hotelId uuid.UUID, lastUpdateTime time.Time) ReservationCanceled {
	s := ReservationCanceled{}
	s.Id = id
	s.HotelId = hotelId
	s.LastUpdateDateTime = lastUpdateTime
	return s
}


