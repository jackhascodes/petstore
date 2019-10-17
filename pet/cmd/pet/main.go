// package main contains the basic setup for the application.
// For the sake of time, a simple environment vars setup has been used to initialise database credentials and api-key.
// Given more time, this would have been done using an env file and go flags to populate these details.
//
// For simplicity's sake, the ListenAndServe functionality is not made concurrent through use of a channel. In a 'real'
// application, consideration would be made to do so depending on frequency and volume of use.
package main

import (

	"github.com/jackhascodes/petstore/pet/cmd/pet/handlers"
	"github.com/jackhascodes/petstore/pet/internal/pet"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()
func main() {

	var (
		host = os.Getenv("dbhost")
		port = os.Getenv("dbport")
		user = os.Getenv("dbuser")
		pass = os.Getenv("dbpass")
		dbName = os.Getenv("dbname")
		apikey = os.Getenv("apikey")
	)
	db, err := pet.InitMySQLPersistence(host, port, user, pass, dbName)
	if err != nil {
		log.Fatalf("error connecting to database, aborting. %v", err)
		os.Exit(2)
	}
	srv := handlers.SetupHttpHandlers(apikey, pet.NewService(db))
	log.Info("pet-app started")
	log.Fatal(srv.ListenAndServe())

}
