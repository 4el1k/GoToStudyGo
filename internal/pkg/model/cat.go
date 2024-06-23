package model

import (
	uuid "github.com/satori/go.uuid"
)

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/model/cat.go

//easyjson:json
type Cat struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	PasswordHash []byte    `json:"-"`
}
