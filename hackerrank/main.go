package main

import (
	"fmt"

	"github.com/xaratustra/hackerrank/coin"
	"github.com/xaratustra/hackerrank/fibo"
	"github.com/xaratustra/hackerrank/paren"
	"github.com/xaratustra/hackerrank/ransom"
)

func main() {

	var option int

	fmt.Println("Which Excersice do you want to run?")
	fmt.Println("1) Ransom Note")
	fmt.Println("2) Coin Change")
	fmt.Println("3) Fibonacci")
	fmt.Println("4) Parenthesis Match")
	fmt.Println("------------------------------------")
	//fmt.Scanln(&option)
	option = 4
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
	}
}
