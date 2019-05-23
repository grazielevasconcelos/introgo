package prefixo

import (
	"strconv"

	"github.com/grazielevasconcelos/introgo/variaveis/pacotes/operadora"
)

//CapitalSP armazena o ddd da capital de sao paulo
var CapitalSP = 11

//PrefixoNomeDaOperadora ddd e nome operadora
var PrefixoNomeDaOperadora = strconv.Itoa(CapitalSP) + " " + operadora.NomeDaOperadora
