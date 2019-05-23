package main

import (
	"encoding/json"
	"fmt"

	"github.com/grazielevasconcelos/introgo/struct_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.SetValor(10)
	fmt.Printf("casa: %+v\r\n", casa)
	casaJSON, _ := json.Marshal(casa)
	fmt.Printf("casa: %+v\r\n", string(casaJSON))
}
