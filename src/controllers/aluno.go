package controllers

import (
    "net/http"

    "api-go-gin/src/models"
    "api-go-gin/src/services"
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
    context.JSON(200, services.FindAll())
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
    aluno, err := services.FindById(context.Params.ByName("id"))
    if aluno.ID == 0 {
        context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
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
    aluno, err := services.FindByCpf(context.Params.ByName("cpf"))
    if aluno.ID == 0 {
        context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
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
        context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return
    }
    message, err := services.SaveAluno(aluno)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return
    }
    context.JSON(http.StatusOK, gin.H{"Success": message})
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
    if err := context.ShouldBindJSON(&aluno); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return
    }
    message, err := services.EditAluno(aluno, id)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return
    }
    context.JSON(http.StatusOK, gin.H{"Success": message})
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
    id := context.Params.ByName("id")
    message, err := services.DeleteAluno(id)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return
    }
    context.JSON(http.StatusOK, gin.H{"Success": message})
}
