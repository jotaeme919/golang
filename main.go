package main

import "fmt"

func main() {
    var a int
	var b int
 	 fmt.Print("Digite um número")
 	 fmt.Scan(&a)
	 fmt.Print("Digite outro número")
	 fmt.Scan(&b)
	 fmt.Println("A soma é: ", a + b)
	 fmt.Println("A subtração é: ", a - b)
	 fmt.Println("A divisão é: ", a / b)
	 fmt.Println("A multiplicação é: ", a * b)
	 fmt.Println("O resto da divisão é: ", a % b)
 	
}