package main

import (
    "api-go-gin/src/database"
    "api-go-gin/src/routes"
)

func main() {
    database.ConectaComBancoDeDados()
    routes.HandleRequests()
}
