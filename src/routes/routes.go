package routes

import (
    "api-go-gin/src/controllers"
    "github.com/gin-gonic/gin"
)

func HandleRequests() {
    r := gin.Default()
    r.GET("/alunos", controllers.ExibeTodosAlunos)
    r.GET("/alunos/:id", controllers.FindById)
    r.GET("/alunos/search/:cpf", controllers.FindByCpf)
    r.POST("/alunos", controllers.SaveAluno)
    r.PUT("/alunos/:id", controllers.EditAluno)
    r.DELETE("/alunos/:id", controllers.DeleteAluno)
    r.Run()
}
