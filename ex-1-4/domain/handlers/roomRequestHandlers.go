package handlers

import (
	"bus/domain/commands"
	"bus/providers/inmemory"
	"bus/requests"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func MakeReservation(rw http.ResponseWriter, req *http.Request, bus inmemory.Bus, roomsAvailable map[string]int, db *sql.DB) error {

	decoder := json.NewDecoder(req.Body)
	var r requests.ReserveRoomRequest
	err := decoder.Decode(&r)
	if err != nil {
		panic(err)
	}
	command := commands.MakeReservation{
		Id: r.Id,
		HotelId: r.HotelId,
		RoomType: r.RoomType,
	}

	row, err := db.Query("SELECT * FROM reservation where roomType=\"" + r.RoomType + "\"")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	count := 0
	for row.Next() {
		count++
	}

	if count >= roomsAvailable[r.RoomType] {
		b, err := json.Marshal("no rooms of this type remaining")
		if err != nil {
			// Handle Error
		}
		rw.Write(b)
		return nil
	}
	if err := bus.SendCommand(&command, db); err != nil {
		log.Printf("%v", err)
		b, err := json.Marshal(err.Error())
		if err != nil {
			// Handle Error
		}
		rw.Write(b)
		return nil
	}
	json.NewEncoder(rw).Encode(r)
	return nil
}

