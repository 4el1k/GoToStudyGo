package req

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/handler/dto/req/couple_token.go

//easyjson:json
type CoupleToken struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
