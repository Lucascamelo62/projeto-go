package main

import "fmt"

func main() {
	nome := "Lucas Camelo"
	versao := 1.1

	fmt.Println("Olá", nome,)
	fmt.Println("Versão do Sistema: ", versao)

	fmt.Println("1- Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
	fmt.Println("Digite um dos comandos acima:")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("A opção digitida foi: ", comando)
}
