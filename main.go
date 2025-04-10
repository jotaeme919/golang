package main

import (
	"fmt" 

)
func main() {
	var saldo int = 20000
	var saque int
	var deposito int
	var opcao string = "deposito || saque"
    fmt.Println("Olá! Qual das opções você deseja? saque, deposito ou sair? ")
	fmt.Scan(&opcao)
	if opcao == "saque" {
		fmt.Println("Quanto deseja sacar?")
		fmt.Scan(&saque)
		saldo -= saque
	} else if opcao == "deposito" {
		fmt.Println("Quanto você deseja depositar? ")
		fmt.Scan(&deposito)
		saldo += deposito
	} else  {
		fmt.Println("Você saiu da sua conta")
    }
	fmt.Println("Saldo atual: ", saldo)
          


	
}

	



