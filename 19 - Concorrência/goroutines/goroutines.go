package main

import "time"

func main() {
	// CONCORRÊNCIA != PARALELISMO
	go escrever("Olá Mundo")
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		println(texto)
		time.Sleep(time.Second)
	}
}
