package main

import (
	"fmt"
	"os"
)

func main() {

	//fmt.Println("O tipo da variável nome é: ", reflect.TypeOf(nome))
	//fmt.Println("O tipo da variável idade é: ", reflect.TypeOf(idade))
	//fmt.Println("O tipo da variável versão é: ", reflect.TypeOf(versao))

	/*
	** Inicio do sistema
	 */

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Comando desconhecido")
	// }
	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {

	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Logs...")

	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)

	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	// Diferentes formas de declarar variavel
	nome := "Maurício"
	var versao float64 = 1.2
	//var idade = 43

	fmt.Print("Olá Sr(a). ", nome)
	//fmt.Println("Idade", idade)
	fmt.Println(", a versão do sistema em uso é: ", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1 - Monitorar sites")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)
	//fmt.Println("O endereço onde a variavel comando foi alocada é: ", &comando)
	fmt.Println("O comando esolhido foi:", comandoLido)

	return comandoLido
}
