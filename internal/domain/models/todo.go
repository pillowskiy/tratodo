package models

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title" validate:"required,min=3,max=32"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"user_id"`
}
