package main

import (
	"fmt"
	"io"
	"strings"

	"bufio"
	"io/ioutil"
	"net/http"
	"os"

	"strconv"
	"time"
)

const monitoramentos = 1
const delay = 2

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {

		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço o comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {

	nome := "AGAMIN"
	versao := 1.2

	fmt.Println("Olá Mr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {

	var comandoLido int

	fmt.Scan(&comandoLido)

	fmt.Println("Comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {

	fmt.Println("1= Iniciar Monitoramento")
	fmt.Println("2= Exibir logs")
	fmt.Println("0= Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento")

	// sites := []string{"https://www.alura.com.br",
	// 	"https://random-status-code.herokuapp.com",
	// 	"https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)

			testaSite(site)
		}

		// time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testaSite(urlSite string) {
	response, err := http.Get(urlSite)

	if err != nil {
		fmt.Println("Aconteceu um erro", err)
	}

	var statusSite bool
	if response.StatusCode == 200 {
		fmt.Println("Site", urlSite, "carregado com sucesso!")
		statusSite = true

	} else {
		fmt.Println("Site", urlSite, "esta com problema. Status Code", response.StatusCode)
		statusSite = false
	}

	registraLog(urlSite, statusSite)
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("listaSites.txt")

	// arquivo, err := ioutil.ReadFile("listaSites.txt")

	if err != nil {
		fmt.Println("Acontoceu um erro", err)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
