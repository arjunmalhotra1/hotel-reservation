package db

const (
	DBNAME     = "hotel-reservation"
	testDBNAME = "hotel-reservation-test"
	DBURI      = "mongodb://127.0.0.1:27017"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
