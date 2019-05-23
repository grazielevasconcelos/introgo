package model

//Imovel é uma struct que armazena dados
type Imovel struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Nome  string `json:"nome"`
	valor int    `json:"valor"`
}

//SetValor setar valor
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}
