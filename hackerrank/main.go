package main

import (
	"fmt"

	"github.com/xaratustra/hackerrank/ransom"
)

func main() {

	var option int

	fmt.Println("Which Excersice do you want to run?")
	fmt.Println("1) Ransom Note")
	fmt.Println("------------------------------------")
	//fmt.Scanln(&option)
	option = 1
	if option == 1 {
		fmt.Println("Running Ransom Note...")
		ransom.Run()
	}
}
