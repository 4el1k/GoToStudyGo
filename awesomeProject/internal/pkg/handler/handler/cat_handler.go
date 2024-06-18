package handler

import (
	"awesomeProject/internal/pkg/handler/dto/req"
	"awesomeProject/internal/pkg/service"
	"fmt"
	"io"
	"net/http"
)

type CatHandler struct {
	service service.CatService
}

func NewCatHandler(service service.CatService) CatHandler {
	return CatHandler{
		service: service,
	}
}

func (h *CatHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error1 from cat handler %w", err)
		return
	}
	defer r.Body.Close()

	dto := &req.CatReq{}
	err = dto.UnmarshalJSON(body)
	if err != nil {
		fmt.Printf("error2 from cat handler %w", err)
		return
	}
	fmt.Printf("%s", dto.Name)
	fmt.Printf(string(body))
	err = h.service.Create(r.Context(), *dto)
	if err != nil {
		fmt.Printf("error3 from cat handler %w", err)
		return
	}
}
