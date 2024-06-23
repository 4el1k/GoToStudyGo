package handler

import (
	"awesomeProject/internal/pkg/handler/dto/req"
	"awesomeProject/internal/pkg/service"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type AuthHandler struct {
	service *service.JwtService
	log     *logrus.Logger
}

func NewAuthHandler(service *service.JwtService, logger *logrus.Logger) AuthHandler {
	return AuthHandler{
		service: service,
		log:     logger,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	h.log.Printf("start signin")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Printf("read body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	in := &req.SignIn{}
	err = in.UnmarshalJSON(body)
	if err != nil {
		h.log.Printf("unmarshal body err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	coupleTokens, err := h.service.Auth(r.Context(), in)
	if err != nil {
		h.log.Printf("auth err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	resp, err := coupleTokens.MarshalJSON()
	if err != nil {
		h.log.Printf("marshal body err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	h.log.Printf("start refresh token")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Printf("read body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	tokens := &req.CoupleToken{}
	err = tokens.UnmarshalJSON(body)
	if err != nil {
		h.log.Printf("unmarshal body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	refreshToken, err := uuid.FromString(tokens.RefreshToken)
	if err != nil {
		h.log.Printf("refresh token err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	jwtToken, err := h.service.RefreshJwtToken(r.Context(), tokens.AccessToken, &refreshToken)
	if err != nil {
		h.log.Printf("auth err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(jwtToken)
	if err != nil {
		return
	}
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	h.log.Printf("start logout")
	// ToDo: поручаю это задание стажеру :))
}
