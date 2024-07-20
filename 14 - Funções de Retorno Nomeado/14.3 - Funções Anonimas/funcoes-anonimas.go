package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, World!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Hello, World!")

	retorno := func(texto string) string {
		return fmt.Sprintf("Tecebido -> %s", texto)
	}("Hello, World!")

	fmt.Println(retorno)
}
