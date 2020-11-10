package main

import (
	"bus/cqrs"
	"bus/domain"
	"bus/domain/commands"
	hotelevents "bus/domain/events"
	"bus/domain/handlers"
	"bus/providers/geteventstore"
	"bus/providers/inmemory"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

func handleRequests() {

	//os.Remove("sqlite-database.db") //

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	db, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer db.Close()
	createTable(db)

	eventStoreClient := geteventstore.CreateEventStoreConnection()
	factory := &cqrs.Factory{}
	bus := inmemory.NewBus(factory, factory)
	commandsHandler := domain.NewCommandsHandler(cqrs.NewAggregateRepository(geteventstore.NewEventStore(eventStoreClient, factory)), db)
	eventsHandler := &domain.EventsHandler{}

	inventoryByRoomType :=  map[string]int{
		"Presidential": 1,
		"King": 1,
		"Queen": 2,
		"Twin": 4,
	}

	bus.RegisterCommandHandler(&commands.MakeReservation{}, commandsHandler)
	bus.RegisterEventHandler(&hotelevents.ReservationMade{}, eventsHandler)
//virtual workshop ex-9 register your CancelReservation command and ReservationCanceled events here

	http.HandleFunc("/makeReservation", func(w http.ResponseWriter, r *http.Request) {
		handlers.MakeReservation(w, r, *bus, inventoryByRoomType, db)
	})

	//virtual workshop ex-9 route the cancelReservation request to the appropriate handler here

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func createTable(db *sql.DB) {
	createReservationsTableSQL := `CREATE TABLE reservation (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,	
		"reservationId" TEXT UNIQUE,	
		"hotelId" TEXT,
		"roomType" TEXT,
		"lastUpdateDateTime" TEXT		
	  );`

	log.Println("Create reservation table...")
	statement, err := db.Prepare(createReservationsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("reservation table created")
}


func main(){
	handleRequests()
}

