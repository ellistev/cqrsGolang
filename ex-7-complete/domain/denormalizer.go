package domain

import (
	"bus/cqrs"
	hotelevents "bus/domain/events"
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"log"
)

type ReservationDenormalizer struct{
	db *sql.DB
}

func NewReservationDenormalizer(db *sql.DB) ReservationDenormalizer {
	return ReservationDenormalizer{db}
}

func ( d *ReservationDenormalizer)DenormalizeEvent(e *cqrs.DomainEvent) error {

	eventFactory := &cqrs.Factory{}
	_, err := eventFactory.GetEventType(e.Event())
	if err != nil {
		return err
	}

	switch v := e.Event().(type){
		case hotelevents.ReservationMade:
			d.OnReservationMade(v)
	}
	return nil
}

func (s *ReservationDenormalizer) OnReservationMade(e hotelevents.ReservationMade) {

	//get the reservation

	// Get Document
	getResult := getReservation(s.db, e.Id)

	if getResult == nil {
		//doesn't exist
		reservation := &Reservation{}
		reservation.Id = e.Id
		reservation.HotelId = e.HotelId
		reservation.LastUpdateDateTime = e.LastUpdateDateTime
		reservation.RoomType = e.RoomType
		insertReservation(s.db, *reservation)

	} else {
		//make changes as necessary
		getResult.Id = e.Id
		getResult.HotelId = e.HotelId
		getResult.LastUpdateDateTime = e.LastUpdateDateTime
		getResult.RoomType = e.RoomType
		getResult.LastUpdateDateTime = e.LastUpdateDateTime

		replaceReservation(s.db, *getResult)
	}
}

func replaceReservation(db *sql.DB, reservation Reservation) {
	log.Println("Inserting reservation record ...")
	replaceReservationSQL := `REPLACE INTO reservation(reservationId, hotelId, roomType, lastUpdateDateTime) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(replaceReservationSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(reservation.Id, reservation.HotelId, reservation.RoomType, reservation.LastUpdateDateTime)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertReservation(db *sql.DB, reservation Reservation) {
	log.Println("Inserting reservation record ...")
	insertReservationSQL := `INSERT INTO reservation(reservationId, hotelId, roomType, lastUpdateDateTime) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertReservationSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(reservation.Id, reservation.HotelId, reservation.RoomType, reservation.LastUpdateDateTime)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getReservation(db *sql.DB, id uuid.UUID) *Reservation{
	log.Println("Inserting reservation record ...")
	selectReservationSQL := `SELECT * FROM reservation where reservationId = "` + id.String() + "\""
	row, err := db.Query(selectReservationSQL)
	if err != nil {
		return nil
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var reservationId string
		var hotelId string
		var roomType string
		row.Scan(&reservationId, &hotelId, &roomType)
		obj := &Reservation{}
		obj.AggregateRoot = cqrs.NewAggregateRoot("Reservation", id, obj.apply)
		obj.Id = id
		obj.HotelId, _ = uuid.FromString(hotelId)
		obj.RoomType = roomType
		return obj
	}
	return nil
}

