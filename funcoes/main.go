package main

import (
	"fmt"

	"github.com/grazielevasconcelos/introgo/funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Println("Resultado da multiplicação: ", resultado)

	fmt.Println("Resultado da soma: ", matematica.Soma(3, 3))

	resultado = matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Println("Resultado da soma func: ", matematica.Soma(3, 3))

	fmt.Println("Resultado da divisao: ", matematica.Calculo(matematica.Divisor, 4, 2))

	resultado, resto := matematica.DivisorComResto(12, 7)
	fmt.Println("Resultado da divisao com resto: ", resultado, resto)
}
