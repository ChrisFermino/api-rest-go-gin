package routes

import (
    _ "api-go-gin/docs"
    "api-go-gin/src/controllers"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
    r := gin.Default()

    api := r.Group("/api")
    {
        alunos := api.Group("/alunos")
        {
            alunos.GET("", controllers.ExibeTodosAlunos)
            alunos.GET(":id", controllers.FindById)
            alunos.GET("/search/:cpf", controllers.FindByCpf)
            alunos.POST("", controllers.SaveAluno)
            alunos.PUT(":id", controllers.EditAluno)
            alunos.DELETE(":id", controllers.DeleteAluno)
        }
    }
    // use ginSwagger middleware to serve the API docs
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}
