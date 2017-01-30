package stack

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	results := []string{}
	stack := []string{}

	for scanner.Scan() {
		process(scanner.Text(), results, stack)
		count = count + 1
		if count == num {
			break
		}
	}

	for _, res := range results {
		fmt.Println(res)
	}
}

func process(instruction string, results []string, stack []string) {

	num := "0"

	texts := strings.Split(instruction, " ")

	if len(texts) < 1 || len(texts) > 2 {
		fmt.Println("Invalid instruction format")
	}

	num = texts[0]

	if _, err := strconv.Atoi(num); err != nil {
		fmt.Println("Invalid Option")
	}

	switch num {
	case "1":
		if len(texts) < 2 {
			fmt.Println("Queue Value Missing")
		} else {
			queue(texts[1], stack)
		}
	case "2":
		dequeue(stack)
	case "3":
		if len(stack) > 0 {
			print(stack[0], results)
		}
	default:
		fmt.Println("Invalid Option")

	}

	fmt.Println(results)
}

func queue(val string, stack []string) {
	fmt.Println("Queue")
	stack = append(stack, val)
	fmt.Println("------")
	fmt.Println(val)
	fmt.Println(stack)
}

func dequeue(stack []string) {
	fmt.Println("Dequeue")
	stack = stack[1:]
	fmt.Println("------")
	fmt.Println(stack)
}

func print(val string, results []string) {
	fmt.Println("Send to Print")
	results = append(results, val)
	fmt.Println("------")
	fmt.Println(val)
	fmt.Println(results)
	// fmt.Println("------")
	// fmt.Println(val)
	// fmt.Println(valtoprint)
}
