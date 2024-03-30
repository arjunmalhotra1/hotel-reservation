package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arjunmalhotra1/hotel-reservation/db"
	"github.com/arjunmalhotra1/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

type BookRoomParams struct {
	FromDate      time.Time `json:"fromDate"`
	TillDate      time.Time `json:"tillDate"`
	NumberPersons int       `json:"numPersons"`
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}
	user, ok := c.Context().Value("user").(*types.User)

	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}
	booking := types.Booking{
		UserID:        user.ID,
		RoomID:        roomID,
		FromDate:      params.FromDate,
		TillDate:      params.TillDate,
		NumberPersons: params.NumberPersons,
	}

	fmt.Printf("%+v \n\n", booking)
	fmt.Printf("%#v \n\n", booking)

	return nil
}
