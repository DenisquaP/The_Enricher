package entity

type Request struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Response struct {
	Request
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
