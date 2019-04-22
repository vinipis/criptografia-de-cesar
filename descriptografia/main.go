package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	conteudo, _ := ioutil.ReadFile("/home/carlos/teste_golang_descriptografia.txt")

	for _, cripytos := range conteudo {

		switch {
		case cripytos == 96:
			cripytos = 32
			fmt.Print(string(cripytos))
		case cripytos == 161:
			cripytos = 97
			fmt.Print(string(cripytos))
		case cripytos == 165:
			cripytos = 101
			fmt.Print(string(cripytos))
		case cripytos == 41:
			cripytos = 105
			fmt.Print(string(cripytos))
		case cripytos == 47:
			cripytos = 111
			fmt.Print(string(cripytos))
		case cripytos == 53:
			cripytos = 117
			fmt.Print(string(cripytos))
		default:
			fmt.Print(string(cripytos - 4))
		}
	}
}
