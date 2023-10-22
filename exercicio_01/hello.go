package main

import (
	"fmt"
)

func main() {

	// Diferentes formas de declarar variavel
	nome := "Maurício"
	var versao float64 = 1.2
	//var idade = 43
	fmt.Print("Olá Sr. ", nome)
	//fmt.Println("Idade", idade)
	fmt.Println(", a versão do sistema em uso é: ", versao)
	fmt.Println("")

	//fmt.Println("O tipo da variável nome é: ", reflect.TypeOf(nome))
	//fmt.Println("O tipo da variável idade é: ", reflect.TypeOf(idade))
	//fmt.Println("O tipo da variável versão é: ", reflect.TypeOf(versao))

	/*
	** Inicio do sistema
	 */

	fmt.Println("1 - Monitorar sites")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O endereço onde a variavel comando foi alocada é: ", &comando)
	fmt.Println("O comando esolhido foi:", comando)

}
