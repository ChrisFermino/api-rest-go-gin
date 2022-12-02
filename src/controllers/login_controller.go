package controllers

import (
    "api-go-gin/src/database"
    "api-go-gin/src/models"
    "api-go-gin/src/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Login(ctx *gin.Context) {
    var p models.Login
    if err := ctx.ShouldBindJSON(&p); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": "cannot bind JSON: " + err.Error()})
        return
    }

    var user models.User
    dbErr := database.DB.Where("email = ?", p.Email).First(&user).Error
    if dbErr != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"Error": "cannot find user"})
        return
    }

    if user.Password != services.SHA256Encoder(p.Password) {
        ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "invalid credentials"})
        return
    }
    token, err := services.NewJWTService().GenerateToken(user.ID)
    if err != nil {
        ctx.JSON(500, gin.H{"Error": err.Error()})
        return
    }
    ctx.JSON(200, gin.H{"token": token})
}
