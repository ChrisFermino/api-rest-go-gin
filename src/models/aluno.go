package models

type Aluno struct {
    ID   int    `json:"id"`
    Nome string `json:"nome"`
    Cpf  string `json:"cpf"`
    Rg   string `json:"rg"`
}
