package controllers

import (
    "api-go-gin/src/models"
    "api-go-gin/src/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SaveUser(context *gin.Context) {
    var u models.User
    if err := context.ShouldBindJSON(&u); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Error": "cannot bind JSON: " + err.Error()})
        return
    }
    msg, err := services.SaveUser(u)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Error": "cannot save User: " + err.Error()})
        return
    }
    context.JSON(http.StatusCreated, gin.H{"Success": msg})
}
