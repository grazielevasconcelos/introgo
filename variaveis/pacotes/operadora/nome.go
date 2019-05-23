package operadora

import (
	"strconv"

	"github.com/grazielevasconcelos/introgo/variaveis/pacotes/prefixo"
)

//NomeDaOperadora armazena a operadora do usuario
var NomeDaOperadora = "Tim"

//PrefixoNomeOperadora ddd e nome operadora
var PrefixoNomeOperadora = strconv.Itoa(prefixo.CapitalSP) + " " + NomeDaOperadora
