package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edwinnduti/natschat/consts"
	"github.com/edwinnduti/natschat/logger"
	"github.com/edwinnduti/natschat/middlewares"
	"github.com/edwinnduti/natschat/natsConn"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	streamName  = "ORDERS"
	subjectName = "ORDERS.received"
)

// init function
func init() {
	// START
	logger.Info.Println("Initializing the application...")

	// load env file (.env by default)
	err := godotenv.Load()
	if err != nil {
		logger.Error.Printf("Error acquiring ENV values: %v\n", err)
	}

	// log for loaded ENV variables
	logger.Success.Println("Env Values Acquired Successfully")
}

// main function
func main(){
	// declare a new jetstreamServer and error
	js := &natsConn.JetStreamContext{}
	var err error
	js.Conn, err = natsConn.JsConnect()
	if err != nil {
		logger.Error.Printf("Error connecting to NATS Jetstream : %v\n", err)
	}

	// if stream does not exist, create a new stream
	err = js.NewStreamAndSubject(streamName, subjectName)
	if err != nil {
		logger.Error.Printf("Error creating new stream and subject : %v\n", err)
	}

	// create subscriptions
	js.Conn.Subscribe(consts.UserCreatedSubject, middlewares.CreateUser(js.Conn.Msg))

	/* log.Print("publishing an order")
	_, err = js.Conn.Publish(subjectName, []byte("one big burger"))
	if err != nil {
		log.Print(err)
	}

	log.Print("attempting to receive order")
	var order []byte
	done := make(chan bool, 1)
	js.Conn.Subscribe(subjectName, func(m *nats.Msg) {
		order = m.Data
		m.Ack()
		done <- true
	})
 */
	select {
	case <-time.After(5 * time.Second):
		log.Fatalf("failed to get order")
	case <-done:
		log.Printf("got order: %q", order)
	}

	// new mux router
	router := mux.NewRouter()

	// port number
	var PORT string = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	// http web server
	server := &http.Server{
		Addr:    fmt.Sprint(":", PORT),
		Handler: router,
	}

	// log info and listen to connection
	logger.Success.Printf("Listening to logs on Port %v...", PORT)
	server.ListenAndServe()
}