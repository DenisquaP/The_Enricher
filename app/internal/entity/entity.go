package entity

type CreateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UpdateRequest struct {
	UserId        int    `json:"user_id"`
	FieldToUpdate string `json:"field_to_update"`
	NewValue      string `json:"new_value"`
}

type DeleteRequest struct {
	UserID int `json:"user_id"`
}

type ResponseErr struct {
	Err error `json:"error"`
}

type ResponseOk struct {
	Message string `json:"message"`
}
