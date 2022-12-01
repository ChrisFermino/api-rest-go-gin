package models

type Aluno struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Nome string `json:"nome"`
    Cpf  string `json:"cpf"`
    Rg   string `json:"rg"`
}
