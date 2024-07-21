package main

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i <= 45; i++ {
		tarefas <- i
	}
	close(tarefas)
	for i := 0; i <= 45; i++ {
		resultado := <-resultados
		println(resultado)
	}
}
