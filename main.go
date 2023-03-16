package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Definindo os caracteres que podem ser usados na senha
	var letras = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numeros = []rune("0123456789")
	var simbolos = []rune("!@#$%&*()_+")

	// Solicitando ao usuário o comprimento da senha desejada
	var comprimento int
	fmt.Println("Digite o comprimento da senha desejada:")
	fmt.Scanln(&comprimento)

	// Solicitando ao usuário a complexidade da senha desejada
	var complexidade int
	fmt.Println("Digite a complexidade da senha desejada (1: apenas letras, 2: letras e números, 3: letras, números e símbolos):")
	fmt.Scanln(&complexidade)

	// Gerando a senha aleatória
	rand.Seed(time.Now().UnixNano())
	var senha = make([]rune, comprimento)
	for i := range senha {
		switch complexidade {
		case 1:
			senha[i] = letras[rand.Intn(len(letras))]
		case 2:
			if rand.Intn(2) == 0 {
				senha[i] = letras[rand.Intn(len(letras))]
			} else {
				senha[i] = numeros[rand.Intn(len(numeros))]
			}
		case 3:
			switch rand.Intn(3) {
			case 0:
				senha[i] = letras[rand.Intn(len(letras))]
			case 1:
				senha[i] = numeros[rand.Intn(len(numeros))]
			case 2:
				senha[i] = simbolos[rand.Intn(len(simbolos))]
			}
		}
	}

	// Exibindo a senha gerada
	fmt.Println("Senha gerada:", string(senha))
}

