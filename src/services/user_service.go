package services

import (
    "api-go-gin/src/database"
    "api-go-gin/src/models"
)

func SaveUser(user models.User) (string, error) {
    //UserDB, err := FindByEmail(aluno.Email)
    //if err != nil {
    //    return "", err
    //} else if UserDB.ID != 0 {
    //    return "", errors.New("there is already a student with this CPF")
    //}
    user.Password = SHA256Encoder(user.Password)
    database.DB.Save(&user)
    return "Student successfully registered", nil
}
