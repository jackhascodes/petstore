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
