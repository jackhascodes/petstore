package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackhascodes/petstore/pet/internal/pet"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)
var log = logrus.New()
func SetupHttpHandlers(k string, s *pet.Service) *http.Server{

	r := mux.NewRouter()
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

func getPetById(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func updatePetForm(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK)

	}
}

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
		w.WriteHeader(http.StatusOK)
	}
}

func uploadImage(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("not implemented"))
	}
}
func addPet(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusOK)

		j, _ := json.Marshal(p)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func updatePet(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusOK)
	}
}

func findByStatus(k string, s *pet.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key != k {
			w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}