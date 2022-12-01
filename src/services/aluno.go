package services

import (
    "api-go-gin/src/database"
    "api-go-gin/src/models"
    "errors"
)

func FindAll() []models.Aluno {
    var alunos []models.Aluno
    database.DB.Order("id").Find(&alunos)
    return alunos
}

func FindById(id string) (models.Aluno, error) {
    var aluno models.Aluno
    if err := database.DB.First(&aluno, id); err != nil {
        return aluno, err.Error
    }
    return aluno, nil
}

func FindByCpf(cpf string) (models.Aluno, error) {
    var aluno models.Aluno
    if err := database.DB.Where("cpf LIKE ?", "%"+cpf+"%").Find(&aluno); err != nil {
        return aluno, err.Error
    }
    return aluno, nil
}

func SaveAluno(aluno models.Aluno) (string, error) {
    alunoBanco, err := FindByCpf(aluno.Cpf)
    if err != nil {
        return "", err
    } else if alunoBanco.ID != 0 {
        return "", errors.New("there is already a student with this CPF")
    }
    database.DB.Save(&aluno)
    return "Student successfully registered", nil
}

func EditAluno(aluno models.Aluno, id string) (string, error) {
    alunoBanco, err := FindById(id)
    if err != nil {
        return "", err
    }
    database.DB.Model(&alunoBanco).UpdateColumns(aluno)
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
