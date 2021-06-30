package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeNomes()
	exibeIntroducao()

	for {
		exibeMenu()

		/*função que retorna mais de uma variável. Continua lá embaixo, aqui só chama a função
		nome, idade := devolveNomeEIdade()

		fmt.Println(nome, "E tenho ", idade, "anos")*/

		comando := leComando()

		//controlando fluxo com switch
		switch comando {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
			//Sai do programa
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
			//Deu erro no programa
		}
	}
}

func exibeIntroducao() {
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitor amento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://alura.com.br",
		"https://www.caelum.com.br"}

	//fmt.Println(sites)

	for i, site := range sites {
		fmt.Println("Estou passando na posicao", i, "do meu slice e essa posiçao tem o site:", site)
	}

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com com problemas. Status Code:",
			resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
