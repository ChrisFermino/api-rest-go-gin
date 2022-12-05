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
func ExibeTodosAlunos(ctx *gin.Context) {
	ctx.JSON(200, services.FindAll())
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
func FindById(ctx *gin.Context) {
	a, err := services.FindById(ctx.Params.ByName("id"))
	if a.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, a)
}

// FindByCpf godoc
// @Summary      Get by Cpf
// @Description  get aluno por CPF
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "Aluno Cpf"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/search/{cpf} [get]
func FindByCpf(ctx *gin.Context) {
	a, err := services.FindByCpf(ctx.Params.ByName("cpf"))
	if a.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, a)
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
func SaveAluno(ctx *gin.Context) {
	var a models.Aluno
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "cannot bind JSON: " + err.Error()})
		return
	}
	msg, err := services.SaveAluno(a)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "cannot save student: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Success": msg})
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
func EditAluno(ctx *gin.Context) {
	var a models.Aluno
	id := ctx.Params.ByName("id")
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	msg, err := services.EditAluno(a, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Success": msg})
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
func DeleteAluno(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	msg, err := services.DeleteAluno(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Success": msg})
}
