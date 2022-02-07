package middlewares

import (
	"encoding/json"

	"github.com/edwinnduti/natschat/db"
	"github.com/edwinnduti/natschat/logger"
	"github.com/edwinnduti/natschat/models"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// send message when done
var (
	createUserDone = make(chan models.UserCreated, 1)
)

func CreateUser(msg *nats.Msg) (string, error) {
	

	// unmarshal incoming data to house
	user, err := json.Marshal(msg.Data)
	if err != nil {
		logger.Error.Println("Marshal error", err)
		return "", err
	}

	// save user to database
	id := uuid.New().String()
	client, err := db.ConnectDb()
	if err != nil {
        logger.Error.Println("Cannot connect to DB", err)
		
    }
	err = client.Set(id, user, 0).Err()
    if err != nil {
        logger.Error.Println("Cannot Add to DB", err)
		return "", err
    }

	// send message when done
	var createdUser models.UserCreated 
	createdUser.ID = id

	// acknowledge message
	msg.Ack()
	createUserDone <- createdUser
}