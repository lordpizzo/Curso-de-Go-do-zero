package main

import (
	"fmt"
	"log"
	"time"
)

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	start := time.Now()
	fmt.Println("Funções Recursivas")
	// resultado := fibonacci(45)
	// fmt.Println(resultado)

	posicao := uint(45)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	elapsed := time.Since(start)
	log.Printf("Recursive funcions took %s", elapsed)
}
