package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/arjunmalhotra1/hotel-reservation/db"
	"github.com/arjunmalhotra1/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27017"
	dbname    = "hotel-reservation-test"
)

type testdb struct {
	db.UserStore
}

func (tdb testdb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	// We could use the in memory implementation of the UserStore for our tests.
	// But if we use the actual Mongo db for tests, we can test our Handler functionality, our Domain functionality
	// and database functionality all in one test. If we test with in-memory implementation we are not testing Mongo implementation.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}

}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "weqeqweqweqweqe",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expecting a user if to be set")
	}

	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expecting the encrypted password not to be included in the Json response")
	}

	if user.FirstName != params.FirstName {
		t.Errorf("expected username%s but got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("expected username%s but got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("expected username%s but got %s", params.Email, user.Email)
	}

	fmt.Println(user)

}
