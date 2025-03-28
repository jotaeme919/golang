package main

import "fmt"

	func main() {
		var usuario string
		var senha string

		fmt.Printf("Digite seu usuário e uma senha:")
        fmt.Scan(&usuario)
        fmt.Scan(&senha)
        if (usuario == "admin" && senha == "1234"){
        fmt.Printf("Usuário correto")
		} else {
			fmt.Println("senha incorreta")
		}
		
        
		
		}