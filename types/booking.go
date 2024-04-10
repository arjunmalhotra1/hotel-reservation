package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID        primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	NumberPersons int                `bson:"numberPersons,omitempty" json:"numberPersons,omitempty"`
	RoomID        primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	FromDate      time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate      time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
	Cancelled     bool               `bson:"cancelled" json:"cancelled,"`
}
