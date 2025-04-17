package main

import (
	"fmt" )

func main() {
	alunoIdade := make(map[string]int)
	alunoIdade["Bruno"] = 16 
	alunoIdade["Otávio"] = 16
	alunoIdade["JM"] = 16
	alunoIdade["Lele"] = 15
	fmt.Println("Idade do Bruno", alunoIdade["Bruno"])
	notasAlunos := map[string]float64{
		"Bruno" : 1.7,
		"Otávio" : 7.5,
		"JM" : 9.8,
		"Lele" : 10,
} 
 for k, v := range notasAlunos{
	fmt.Printf("%s tirou a nota %.1f \n", k, v)

	}


 

}
