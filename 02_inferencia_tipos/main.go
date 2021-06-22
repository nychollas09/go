package main

import (
	"fmt"
	"reflect"
)

// O ":" serve para declarar ao mesmo tempo que atribui um valor a variável

func main() {
	nome := "Nichollas"
	idade := 20
	versao := 1.1
	fmt.Println("Olá, sr.", nome, ", minha idade é ", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é: ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é: ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versão é: ", reflect.TypeOf(versao))
}
