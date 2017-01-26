package paren

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Run excersise
func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	results := make([]string, num)

	for scanner.Scan() {
		results[count] = isParenBalanced(scanner.Text())
		count = count + 1
		if count == num {
			break
		}
	}

	for _, res := range results {
		fmt.Println(res)
	}

}

func isParenBalanced(line string) string {

	NO := "NO"
	YES := "YES"

	result := YES

	if len(line) == 0 {
		return result
	}

	if len(line)%2 != 0 {
		return NO
	}

	stack := []rune{}
	validParen := []rune{'(', ')', '{', '}', '[', ']'}

	for _, r := range line {
		if !contains(r, validParen) {
			return NO
		}

		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else if len(stack) > 0 && ((r == ')' && stack[len(stack)-1] == '(') || (r == '}' && stack[len(stack)-1] == '{') || (r == ']' && stack[len(stack)-1] == '[')) {
			stack = stack[:len(stack)-1]
		} else {
			return NO
		}
	}

	if len(stack) != 0 {
		return NO
	}

	return result
}

func contains(r rune, list []rune) bool {
	for _, val := range list {
		if val == r {
			return true
		}
	}

	return false
}
