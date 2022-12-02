package services

import (
    "api-go-gin/src/database"
    "api-go-gin/src/models"
    "errors"
)

func FindAll() []models.Aluno {
    var a []models.Aluno
    database.DB.Order("id").Find(&a)
    return a
}

func FindById(id string) (models.Aluno, error) {
    var a models.Aluno
    if err := database.DB.First(&a, id); err != nil {
        return a, err.Error
    }
    return a, nil
}

func FindByCpf(cpf string) (models.Aluno, error) {
    var a models.Aluno
    if err := database.DB.Where("cpf LIKE ?", "%"+cpf+"%").First(&a); err != nil {
        return a, err.Error
    }
    return a, nil
}

func SaveAluno(a models.Aluno) (string, error) {
    aDB, _ := FindByCpf(a.Cpf)
    if aDB.ID != 0 {
        return "", errors.New("there is already a student with this CPF")
    }
    database.DB.Save(&a)
    return "Student successfully registered", nil
}

func EditAluno(a models.Aluno, id string) (string, error) {
    aDB, err := FindById(id)
    if err != nil {
        return "", err
    }
    database.DB.Model(&aDB).UpdateColumns(a)
    return "Student updated successfully", nil
}

func DeleteAluno(id string) (string, error) {
    _, err := FindById(id)
    if err != nil {
        return "", err
    }
    database.DB.Delete(&models.Aluno{}, id)
    return "Student updated successfully", nil
}
