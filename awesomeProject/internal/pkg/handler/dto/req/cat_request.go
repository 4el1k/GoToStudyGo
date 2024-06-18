package req

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/handler/dto/req/cat_request.go

//easyjson:json
type CatReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
