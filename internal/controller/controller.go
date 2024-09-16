package controller

import (
	"encoding/json"
	"net/http"
	"repo/internal/repository"
	"repo/internal/service"
	"strconv"
)

type Transport struct {
	UserService service.Service
}

func NewTransport(svc service.Service) *Transport {
	return &Transport{
		UserService: svc,
	}
}

func (t Transport) Create(w http.ResponseWriter, r *http.Request) {
	var user repository.User

	json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	t.UserService.ServiceCreate(user)
}

func (t Transport) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	name := t.UserService.ServiceGetByID(int64(i))

	json.NewEncoder(w).Encode(name)
	w.WriteHeader(http.StatusOK)
}

func (t Transport) Update(w http.ResponseWriter, r *http.Request) {
	var user repository.User

	json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	err := t.UserService.ServiceUpdate(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (t Transport) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	t.UserService.ServiceDelete(int64(i))
}

func (t Transport) List(w http.ResponseWriter, r *http.Request) {
	users, err := t.UserService.ServiceList()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}
