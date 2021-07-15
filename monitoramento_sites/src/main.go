package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const totalAttemptsForMonitoring = 5
const delayMonitoring = 5

var clear map[string]func() = make(map[string]func())

func defineCleanStrategy() {
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Ocorreu um problema: ", err.Error())
	}
}

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

func catchSitesFromFile() (sites []string) {
	file, err := os.Open("sites.txt")

	errorHandler(err)

	reader := bufio.NewReader(file)

	var sitesFromFile []string

	for {
		row, err := reader.ReadString('\n')
		sitesFromFile = append(sitesFromFile, strings.TrimSpace(row))

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sitesFromFile
}

func registerLog(url string, status bool) {
	file, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	errorHandler(err)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - status: " + func() string {
		if status {
			return "online"
		} else {
			return "offline"
		}
	}() + "\n")

	file.Close()
}

func testingTarget(url string) {
	response, err := http.Get(url)

	errorHandler(err)

	if response.StatusCode == 200 {
		fmt.Println("Site: ", url, "foi carregado com sucesso!")
		registerLog(url, true)
	} else {
		fmt.Println("Site: ", url, "está com problemas. Status Code:", response.StatusCode)
		registerLog(url, false)
	}
}

func showLogs() {
	file, err := ioutil.ReadFile("logs.log")

	errorHandler(err)

	fmt.Println(string(file))
}

func startMonitoring() {
	sites := catchSitesFromFile()

	for i := 0; i < totalAttemptsForMonitoring; i++ {
		for _, url := range sites {
			testingTarget(url)
		}
		time.Sleep(delayMonitoring * time.Second)
		fmt.Println()
	}
}

func main() {
	defineCleanStrategy()
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
			callClear()
		case 2:
			showLogs()
		default:
			fmt.Println("Não conheço esse comando...")
			os.Exit(-1)
		}
	}
}
