package matematica

/*
Calculo Realiza o calculo matematico de dois numeros inteiros
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica dois numeros
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor divide dois numeros
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto divide dois numeros
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	// if resultado > 1 {
	// 	fmt.Println("aqui")
	// }

	// if resultado > 0 {
	// 	fmt.Println("aqui")
	// 	return
	// }

	resto = numeroA % numeroB

	return
}
