package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const totalAttemptsForMonitoring = 5
const delayMonitoring = 5

func introduction() {
	nome := "Nichollas"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readComand() (comand int) {
	fmt.Println("Escolha um comando")

	// O "&" extrai o endereço/ponteiro da memória da variável informada
	var comando int
	fmt.Scan(&comando)

	return comando
}

func testingTarget(url string) {
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		fmt.Println("Site: ", url, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", url, "está com problemas. Status Code:", response.StatusCode)
	}
}

func startMonitoring() {
	sites := []string{"https://random-status-code.herokuapp.com", "https://alura.com.br", "https://caelum.com.br"}

	for i := 0; i < totalAttemptsForMonitoring; i++ {
		for _, url := range sites {
			testingTarget(url)
		}
		time.Sleep(delayMonitoring * time.Second)
		fmt.Println()
	}
}

func main() {
	introduction()
	fmt.Println()

	for {
		showMenu()
		fmt.Println()

		comand := readComand()
		fmt.Println()

		switch comand {
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		default:
			fmt.Println("Não conheço esse comando...")
			os.Exit(-1)
		}
	}
}
