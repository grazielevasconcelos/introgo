package main

import (
	"encoding/json"
	"fmt"
)

//Imovel é uma struct que armazena dados
type Imovel struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Nome  string `json:"nome"`
	Valor int    `json:"valor"`
}

func main() {
	casa := Imovel{}
	fmt.Println("a casa é: ", casa)

	ap := Imovel{17, 15, "meu ap", 100000}
	fmt.Println("ap é: ", ap)
	apJSON, _ := json.Marshal(ap)
	fmt.Println("JSON: ", string(apJSON))

	fazenda := Imovel{
		Y:     85,
		Nome:  "Chacara",
		Valor: 55,
		X:     22,
	}
	fmt.Println("fazenda é: ", fazenda)

	casa.Nome = "casa"
	fmt.Println("a casa é: ", casa)

	lar := model.Imovel
	lar.Valor = 55
	fmt.Println("casa: ", casa)

}
