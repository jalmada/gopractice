package fibo

import (
	"fmt"
	"time"
)

func Run() {
	var n int
	fmt.Scanf("%d\n", &n)

	start := time.Now()
	fmt.Println(fibonacci(n))
	fmt.Println(time.Since(start))
}

func fibonacci(pos int) int {
	if pos == 0 {
		return 0
	}

	if pos == 1 {
		return 1
	}

	value := 0

	fibo1 := fibonacci(pos - 2)
	fibo2 := fibonacci(pos - 1)
	value = fibo1 + fibo2

	return value
}
