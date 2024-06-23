package req

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/handler/dto/req/sign_in.go

//easyjson:json
type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
