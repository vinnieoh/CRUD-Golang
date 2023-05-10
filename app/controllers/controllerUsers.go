package controllers

import (
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func DataBateConnct(db *gorm.DB) handler {
	return handler{db}
}

func CreateUser() {

}

func GetAllUser() {

}

func GetOneUser() {

}

func PostUser() {

}

func DeleteUser() {

}
