package controllers

import (
    "api-go-gin/src/database"
    "api-go-gin/src/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func ExibeTodosAlunos(context *gin.Context) {
    var alunos []models.Aluno
    database.DB.Find(&alunos)
    context.JSON(200, alunos)
}

func FindById(context *gin.Context) {
    var aluno models.Aluno
    id := context.Params.ByName("id")
    database.DB.First(&aluno, id)
    if aluno.ID == 0 {
        context.JSON(http.StatusNotFound, gin.H{
            "Not Found": "aluno não encontrado"})
        return
    }
    context.JSON(http.StatusOK, aluno)
}

func FindByCpf(context *gin.Context) {
    var aluno models.Aluno
    cpf := context.Param("cpf")
    database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno)
    if aluno.ID == 0 {
        context.JSON(http.StatusNotFound, gin.H{
            "Not Found": "aluno não encontrado"})
        return
    }
    context.JSON(http.StatusOK, aluno)
}

func SaveAluno(context *gin.Context) {
    var aluno models.Aluno
    if err := context.ShouldBindJSON(&aluno); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&aluno)
    context.JSON(http.StatusOK, aluno)
}

func EditAluno(context *gin.Context) {
    var aluno models.Aluno
    id := context.Params.ByName("id")
    database.DB.First(&aluno, id)
    if err := context.ShouldBindJSON(&aluno); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Model(&aluno).UpdateColumns(aluno)
    context.JSON(http.StatusOK, aluno)
}

func DeleteAluno(context *gin.Context) {
    var aluno models.Aluno
    id := context.Params.ByName("id")
    database.DB.Delete(&aluno, id)
    context.JSON(http.StatusOK, gin.H{"data": "aluno deletado com sucesso"})
}
