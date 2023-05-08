package models

type User struct {
	ID       int64  `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
typer User struct{
	ID int64 `json:"id"`
	Nome string `json:"nome"`
	Telefone string `json:"telefone"`
	Birthday Date
	Email string `json:"email"`
	Password string `json:"password"`
	createdAt
  	updatedAt
}
*/
