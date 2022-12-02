package models

type Aluno struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Cpf  string `json:"cpf"`
    Rg   string `json:"rg"`
}
