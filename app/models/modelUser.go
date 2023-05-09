package models

type User struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Birthday  string `json:"birthday"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
