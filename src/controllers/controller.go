package controllers

import (
    "net/http"

    "api-go-gin/src/database"
    "api-go-gin/src/models"
    "github.com/gin-gonic/gin"

    _ "api-go-gin/docs"
    _ "api-go-gin/src/httputil"
)

// ExibeTodosAlunos godoc
// @Summary      Get All
// @Description  Exibe todos os alunos
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos [get]
func ExibeTodosAlunos(context *gin.Context) {
    var alunos []models.Aluno
    database.DB.Find(&alunos)
    context.JSON(200, alunos)
}

// FindById godoc
// @Summary      Get by ID
// @Description  Busca um aluno por ID
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Aluno ID"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{id} [get]
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

// FindByCpf godoc
// @Summary      Get by Cpf
// @Description  get aluno por CPF
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "Aluno Cpf"
// @Success      200  {object}  models.Aluno
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/search/{cpf} [get]
func FindByCpf(context *gin.Context) {
    var aluno models.Aluno
    cpf := context.Param("cpf")
    database.DB.Where("cpf LIKE ?", "%"+cpf+"%").Find(&aluno)
    if aluno.ID == 0 {
        context.JSON(http.StatusNotFound, gin.H{
            "Not Found": "aluno não encontrado"})
        return
    }
    context.JSON(http.StatusOK, aluno)
}

// SaveAluno godoc
// @Summary      Save aluno
// @Description  salva um aluno na tabela alunos
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        aluno   body   models.Aluno   true  "Modelo do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /alunos [post]
func SaveAluno(context *gin.Context) {
    var aluno models.Aluno
    if err := context.ShouldBindJSON(&aluno); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&aluno)
    context.JSON(http.StatusOK, aluno)
}

// EditAluno godoc
// @Summary      Edit aluno
// @Description  edita um aluno na tabela alunos
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        aluno   body   models.Aluno   true  "Modelo do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{id} [put]
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

// DeleteAluno godoc
// @Summary      Delete aluno
// @Description  deleta um aluno no banco de dados
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "id aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{id} [delete]
func DeleteAluno(context *gin.Context) {
    var aluno models.Aluno
    id := context.Params.ByName("id")
    database.DB.Delete(&aluno, id)
    context.JSON(http.StatusOK, gin.H{"data": "aluno deletado com sucesso"})
}
