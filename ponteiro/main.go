package main

import "fmt"

//Imovel é uma struct que armazena dados
type Imovel struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Nome  string `json:"nome"`
	Valor int    `json:"valor"`
}

func main() {
	casa := new(Imovel)
	casa.X = 5
	casa.Y = 10

	fmt.Printf("casa: %p - %+v\r\n", casa, *casa)
	fmt.Println("casa: ", casa, *casa)

	mudaImovel(casa)

	fmt.Printf("casa: %p - %+v\r\n", casa, *casa)
}

func mudaImovel(imovel *Imovel) {
	imovel.Nome = "CASA"
	imovel.Valor = 1000000
}
