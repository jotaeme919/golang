package main

import ("fmt" 

)
func main() {
	var number = [5]int{}
	fmt.Println("Escolha o primeiro número: ")
	fmt.Scan(&number[0])
	fmt.Println("Escolha o segundo número: ")
	fmt.Scan(&number[1])
	fmt.Println("Escolha o terceiro número: ")
	fmt.Scan(&number[2])
	fmt.Println("Escolha o quarto número: ")
	fmt.Scan(&number[3])
	fmt.Println("Escolha o quinto número: ")
	fmt.Scan(&number[4])
	var soma = number[0] + number [1] + number[2] + number[3] + number[4]
	fmt.Println("A soma é: ", soma)


  
}
