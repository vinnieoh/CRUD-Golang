package main

import (
	"api/app/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run()

}