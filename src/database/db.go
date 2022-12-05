package database

import (
	"api-go-gin/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return os.Getenv(key)
}

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=" + envVariable("POSTGRES_USER") + " password=" + envVariable("POSTGRES_PASSWORD") + " dbname=" + envVariable("POSTGRES_DB") + " port=" + envVariable("POSTGRES_PORT") + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
	DB.AutoMigrate(&models.User{})
}
