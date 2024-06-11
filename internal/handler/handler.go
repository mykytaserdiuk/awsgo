package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk/aws-go/internal"
	"github.com/mykytaserdiuk/aws-go/pkg/models"
)

type Handler struct {
	service internal.Service
}

func NewHandler(router *mux.Router, service internal.Service) *Handler {
	h := &Handler{service}
	router.HandleFunc("/", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/", h.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/", h.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/{id}", h.Update).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.GetByID).Methods(http.MethodGet)

	return h
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var request models.TodoIn
	err := request.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	id, err := h.service.Create(r.Context(), request.Topic, request.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var request models.TodoIn
	err := request.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = h.service.Update(r.Context(), id, request.Topic, request.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	err := h.service.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	todo, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
