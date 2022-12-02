package database

import (
    "api-go-gin/src/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var (
    DB  *gorm.DB
    err error
)

func ConectaComBancoDeDados() {
    stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(stringDeConexao))
    if err != nil {
        log.Panic("Erro ao conectar com banco de dados")
    }
    DB.AutoMigrate(&models.Aluno{})
    DB.AutoMigrate(&models.User{})
}
