package req

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/handler/dto/req/refresh_token.go

//easyjson:json
type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
