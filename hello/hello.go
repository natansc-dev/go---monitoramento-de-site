package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5 // Seconds

func main() {
	intro()

	for {
		showOptions()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs...")
			printLogs()
		case 0:
			fmt.Println("Logout...")
			os.Exit(0)
		default:
			fmt.Println("Commands not exist")
			os.Exit(-1)
		}
	}
}

func intro() {
	name := "Natan"
	version := 1.1

	fmt.Println("Hi, Mr.", name)
	fmt.Println("Este programa está na versão", version)
	fmt.Println("")
}

func showOptions() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The command inserted is:", command)
	fmt.Println("")

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	// sites := []string{"https://tbrweb.com.br", "https://abge.org.br"}
	sites := readArchiveSite()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas!")
		registerLog(site, false)
	}
}

func readArchiveSite() []string {
	var sites []string

	doc, err := os.Open("sites.txt")
	// doc, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ops... ocorreu um erro:", err)
	}

	reader := bufio.NewReader(doc)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace((line))

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	doc.Close()

	return sites
}

func registerLog(site string, isOnline bool) {
	doc, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	doc.WriteString("Is online: " + strconv.FormatBool(isOnline) + " - " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + "\n")

	doc.Close()
}

func printLogs() {
	doc, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(doc))
}
