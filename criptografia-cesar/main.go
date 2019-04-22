package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ascii := r1.Intn(60)
	return ascii
}

func main() {

	conteudo, _ := ioutil.ReadFile("/home/carlos/teste_golang_leitura.txt")

	for _, cripytos := range conteudo {

		ascii := numeroAleatorio()

		switch {
		case cripytos == 32 && ascii >= 51:
			cripytos = 96
			fmt.Print(string(cripytos))
		case cripytos == 97 && (ascii >= 41 && ascii <= 50):
			cripytos = 161
			fmt.Print(string(cripytos))
		case cripytos == 101 && (ascii >= 31 && ascii <= 40):
			cripytos = 165
			fmt.Print(string(cripytos))
		case cripytos == 105 && (ascii >= 21 && ascii <= 30):
			cripytos = 41
			fmt.Print(string(cripytos))
		case cripytos == 111 && (ascii >= 11 && ascii <= 20):
			cripytos = 47
			fmt.Print(string(cripytos))
		case cripytos == 117 && (ascii <= 10):
			cripytos = 53
			fmt.Print(string(cripytos))
		default:
			fmt.Print(string(cripytos + 4))
		}
	}
}
