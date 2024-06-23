package handler

import (
	"awesomeProject/internal/pkg/handler/dto/req"
	"awesomeProject/internal/pkg/service"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

type CatHandler struct {
	service *service.CatService
	log     *logrus.Logger
}

func NewCatHandler(service *service.CatService, logger *logrus.Logger) CatHandler {
	return CatHandler{
		service: service,
		log:     logger,
	}
}

func (h *CatHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h.log.Println("start cat handler create")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error1 from cat handler %w", err)
		return
	}
	defer r.Body.Close()
	dto := &req.CatSave{}
	err = dto.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error2 from cat handler %w", err)
		return
	}
	err = h.service.Create(r.Context(), dto)
	if err != nil {
		h.log.Printf("error3 from cat handler %w", err)
		http.Error(w, "tech work", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CatHandler) GetById(w http.ResponseWriter, r *http.Request) {
	h.log.Println("start cat handler get by id")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	h.log.Println("start cat handler get by id", idStr)
	if !ok || idStr == "" {
		h.log.Println("error1 from cat handler get by id")
		http.Error(w, "error1 from cat handler get by id", http.StatusBadRequest)
		return
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		h.log.Println("error2 from cat handler get by id")
		http.Error(w, "error2 from cat handler get by id", http.StatusBadRequest)
		return
	}
	p, err := h.service.FindById(r.Context(), &id)
	resp, err := p.MarshalJSON()
	if err != nil {
		h.log.Println("error3 from cat handler get by id")
		http.Error(w, "tech work", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

func (h *CatHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h.log.Println("start cat handler update")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error1 from cat handler %w", err)
		return
	}
	defer r.Body.Close()
	dto := &req.CatUpdate{}
	err = dto.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error2 from cat handler %w", err)
		return
	}
	err = h.service.Update(r.Context(), dto)
	if err != nil {
		log.Printf("error3 from cat handler %w", err)
		http.Error(w, "tech work", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CatHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	h.log.Println("start cat handler get by id", idStr)
	if !ok || idStr == "" {
		h.log.Println("error1 from cat handler get by id")
		http.Error(w, "error1 from cat handler get by id", http.StatusBadRequest)
		return
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		h.log.Println("error2 from cat handler get by id")
		http.Error(w, "error2 from cat handler get by id", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteById(r.Context(), &id)
	if err != nil {
		h.log.Println("error3 from cat handler get by id")
		http.Error(w, "tech work", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
