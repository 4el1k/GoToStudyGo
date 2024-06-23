package req

//go:generate easyjson -all /home/bebra/GolandProjects/awesomeProject/internal/pkg/handler/dto/req/cat_save.go

//easyjson:json
type CatSave struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Password string `json:"password"`
}
