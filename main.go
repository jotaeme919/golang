package main

import (
	"fmt" )

func DadosPessoais(nome string, idade int) (int, string){
	var classificação string
	if idade >= 18 {
	  classificação = "Você é maior de idade"
	} else {
		classificação = "Você é menor de idade"
}
return idade, classificação

}

func main() {
	var idadepessoa int
	var nomepessoa string
	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nomepessoa)
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&idadepessoa)
	idade, classificação := DadosPessoais(nomepessoa, idadepessoa)
	fmt.Println(idade)
	fmt.Println(classificação)
	
 } 



 


