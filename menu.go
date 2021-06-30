package main

import "fmt"

//func main() {
	nome := "Guilherme"
	reflect.typeOf(nome)
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	//"%d" representa um inteiro e o que o usuário digitar gravar na variavel comando
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	//Controle de fluxo com if
	/*
		if comando == 1 {
			fmt.Println("Monitorando...")
		} else if comando == 2 {
			fmt.Println("Exibindo Logs...")
		} else {
			fmt.Println("Não conheço este comando")
		}*/

	//controlando fluxo com switch
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")
	}
}
