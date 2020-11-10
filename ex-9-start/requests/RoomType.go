package requests

import (
	"encoding/json"
	"errors"
)

type RoomType string

const(
	Presidential RoomType = "Presidential"
	King = "King"
	Queen = "Queen"
	Twin = "Twin"
)

func (lt *RoomType) UnmarshalJSON(b []byte) error {
	type LT RoomType
	var r *LT = (*LT)(lt)
	err := json.Unmarshal(b, &r)
	if err != nil{
		panic(err)
	}
	switch *lt {
	case Presidential, King, Queen, Twin:
		return nil
	}
	return errors.New("invalid room type")
}

func (lt RoomType) IsValid() error {
	switch lt {
	case Presidential, King, Queen, Twin:
		return nil
	}
	return errors.New("invalid room type")
}