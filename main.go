package main

import (
    "api-go-gin/src/database"
    "api-go-gin/src/routes"

    _ "api-go-gin/docs"
    _ "api-go-gin/src/controllers"
    _ "api-go-gin/src/httputil"
    _ "api-go-gin/src/models"
    _ "gorm.io/gorm"
)

// @title           API rest go-gin
// @version         1.0
// @description     This is a REST API with gin Framework
// @host            localhost:8080
// @BasePath        /api
func main() {
    database.ConectaComBancoDeDados()
    routes.HandleRequests()
}
