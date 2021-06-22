package main

import (
	"fmt"
)

func main() {
	nome := "Nichollas"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	// O "&" extrai o endereço/ponteiro da memória da variável informada
	var comando int
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	}else if comando ==2 {
		fmt.Println("Exibindo logs...")
	}else if comando == 0 {
		fmt.Println("Saindo do programa...")
	}else{
		fmt.Println("Não conheço este comando")
	}
}
