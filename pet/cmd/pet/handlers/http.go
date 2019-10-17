package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackhascodes/petstore/pet/internal/pet"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SetupHttpHandlers(k string, s *pet.Service) *http.Server {

	r := mux.NewRouter()

	// Handlers are returned via wrapped funcs to allow service sharing.
	r.HandleFunc("/pet/{id}", getPetById(k, s)).Methods("GET")
	r.HandleFunc("/pet/{id}", updatePetForm(k, s)).Methods("POST")
	r.HandleFunc("/pet/{id}", deletePet(k, s)).Methods("DELETE")
	r.HandleFunc("/pet/{id}/uploadImage", uploadImage(k, s)).Methods("POST")
	r.HandleFunc("/pet", addPet(k, s)).Methods("POST")
	r.HandleFunc("/pet", updatePet(k, s)).Methods("PUT")
	r.HandleFunc("/pet/findByStatus", findByStatus(k, s)).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}

// getPetById returns a handler which will perform very basic auth and responds with:
// * 401 Unauthorized - if the api key is incorrect
// * 404 Unfound - if the requested pet does not exist
// * 200 - and a json representation of the pet if the pet is found
func getPetById(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res := s.FindById(id)
		if res == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, _ := json.Marshal(res)
		w.Write(j)
	}
}

// updatePetForm returns a handler which converts form data to a Pet instance and tries an update. Returns:
// 401 Unauthorized - if the api key is incorrect or missing
// 400 Bad Request - if the input cannot be parsed correctly
// 200 OK - if the update happened without error
func updatePetForm(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.ParseForm()
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p := pet.InitPet(r.FormValue("name"), strings.Split(r.FormValue("photoUrls"), ","))
		p.Id = id
		err = s.UpdatePet(p)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}

// deletePet returns a handler which soft-deletes a Pet in the database. Returns:
// 401 Unauthorized - if the api key is incorrect or missing
// 400 Bad Request - if the id cannot be parsed correctly
// 200 OK - if the delete happened without error
func deletePet(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = s.DeletePet(id)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

// uploadImage has not been implemented and will tell you so.
func uploadImage(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("not implemented"))
	}
}

// addPet takes a json representation of a pet and inserts it into the database. Returns:
// 401 Unauthorized - if the api key is missing or incorrect
// 400 Bad Request - if the json cannot be decoded
// 200 OK - and a json representation of the pet with its new id if the insert happened with no errors
func addPet(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		p := &pet.Pet{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p, err = s.AddPet(p)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		j, _ := json.Marshal(p)
		w.Write(j)
	}
}

// updatePet takes a json representation of the Pet struct and updates an existing instance in the database. Returns:
// 401 Unauthorized - if the api key is missing or incorrect
// 400 Bad Request - if there are any issues decoding the json
// 200 OK - if the update was all good.
func updatePet(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		p := &pet.Pet{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = s.UpdatePet(p)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// findByStatus returns an array of pets with the relevant status. Returns:
// 401 Unauthorized - if the api key is missing or incorrect
// 200 OK - and a json array of pet objects
func findByStatus(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		st := pet.INVALID
		err := st.FromString(vars["status"])
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res := s.FindByStatus(st)

		j, _ := json.Marshal(res)
		w.Write(j)
	}
}
