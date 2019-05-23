package main

import (
	"fmt"

	"github.com/grazielevasconcelos/introgo/variaveis/pacotes/operadora"
	"github.com/grazielevasconcelos/introgo/variaveis/pacotes/prefixo"
)

//NomeDoUsuario armazena o nome do usuario do sistema
var NomeDoUsuario = "Grazi"

func main() {
	fmt.Printf("Nome do usuario do sistema: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital paulista: %d\r\n", prefixo.CapitalSP)
	fmt.Printf("Nome da Operadora do usuario: %s\r\n", operadora.NomeDaOperadora)
}
