package req

import uuid "github.com/satori/go.uuid"

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/handler/dto/req/cat_update.go

//easyjson:json
type CatUpdate struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Id   uuid.UUID `json:"id"`
}
