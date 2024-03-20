package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arjunmalhotra1/hotel-reservation/db"
	"github.com/arjunmalhotra1/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	HotelStore := db.NewMongoHotelStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Bellucia",
		Location: "France",
	}

	room := types.Room{
		Type:      types.SinglePersonRoomType,
		BasePrice: 99.9,
	}
	_ = room

	insertedHotel, err := HotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedHotel)
}
