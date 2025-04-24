package main

import (
	"fmt" )

func pegarNome() (string, string){
	return "Almofa", "dinha"

}
func main() {
	nome, sobrenome := pegarNome()
    fmt.Println(nome)
	fmt.Println(sobrenome)

 

}
