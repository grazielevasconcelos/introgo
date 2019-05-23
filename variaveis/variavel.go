package main

import "fmt"

var (
	//Endereco armazena o endereco do usuario
	Endereco string
	//Telefone armazena o telefone do usuario
	Telefone   = "11111-99999"
	celular    string  // ""
	quantidade int     //0
	comprou    bool    //false
	valor      float64 //0.00
	palavras   rune
)

func main() {
	teste := "Valor de teste"

	fmt.Println("endereco: ", Endereco)
	fmt.Println("telefone: ", Telefone)
	fmt.Println("quantidade: ", quantidade)
	fmt.Println("comprou: ", comprou)
	fmt.Println("valor: ", valor)
	fmt.Println("palavras: ", palavras)
	fmt.Println("teste: ", teste)
}
