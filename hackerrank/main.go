package main

import (
	"fmt"

	"github.com/xaratustra/hackerrank/coin"
	"github.com/xaratustra/hackerrank/fibo"
	"github.com/xaratustra/hackerrank/median"
	"github.com/xaratustra/hackerrank/paren"
	"github.com/xaratustra/hackerrank/ransom"
	"github.com/xaratustra/hackerrank/stack"
)

func main() {

	var option int

	fmt.Println("Which Excersice do you want to run?")
	fmt.Println("1) Ransom Note")
	fmt.Println("2) Coin Change")
	fmt.Println("3) Fibonacci")
	fmt.Println("4) Parenthesis Match")
	fmt.Println("5) Two Stacks")
	fmt.Println("6) Get the Median with two heaps")
	fmt.Println("------------------------------------")
	//fmt.Scanln(&option)
	option = 6
	if option == 1 {
		fmt.Println("Running Ransom Note...")
		ransom.Run()
	} else if option == 2 {
		fmt.Println("Running Coin Change...")
		coin.Run()
	} else if option == 3 {
		fmt.Println("Running Fibonacci...")
		fibo.Run()
	} else if option == 4 {
		fmt.Println("Running Parenthesis Match...")
		paren.Run()
	} else if option == 5 {
		fmt.Println("Running Stacks...")
		stack.Run()
	} else if option == 6 {
		fmt.Println("Running Median...")
		median.Run()
	}
}
