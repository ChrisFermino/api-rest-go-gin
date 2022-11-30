package services

import (
    "api-go-gin/src/database"
    "api-go-gin/src/models"
)

func FindAll() []models.Aluno {
    var alunos []models.Aluno
    database.DB.Find(&alunos)
    return alunos
}

func FindById(id string) models.Aluno {
    var aluno models.Aluno
    database.DB.First(&aluno, id)
    return aluno
}

func FindByCpf(cpf string) models.Aluno {
    var aluno models.Aluno
    database.DB.Where("cpf LIKE ?", "%"+cpf+"%").Find(&aluno)
    return aluno
}

func SaveAluno(aluno models.Aluno) {
    if alunoBanco := FindByCpf(aluno.Cpf); alunoBanco.ID != 0 {
        return // implementar
    }
    // implementar
}

func EditAluno(aluno models.Aluno, id string) {
    // Implementar
}

func DeleteAluno(id string) {
    // Implementar
}
