package coin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Run Coin  Change
func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	var value, noValues int

	scanner.Scan()
	text := scanner.Text()
	values := strings.Split(text, " ")
	value, _ = strconv.Atoi(values[0])
	noValues, _ = strconv.Atoi(values[1])

	scanner.Scan()
	coinValuesText := scanner.Text()

	coinValues := make([]int, noValues, noValues)
	for i, cv := range strings.Split(coinValuesText, " ") {
		if cv != "" && cv != " " {
			coinValues[i], _ = strconv.Atoi(cv)
		}
	}

	operations := make([][]int, value+1)
	for i := 0; i < value+1; i++ {
		operations[i] = make([]int, noValues)
		for j := 0; j < noValues; j++ {

			operations[i][j] = -1
		}
	}

	for i := 0; i < noValues; i++ {
		operations[0][i] = 1
	}

	result := calc(coinValues, noValues, value, operations) //, "center", 0, "")

	fmt.Println(result)
}

func calc(coinValues []int, noValues int, value int, operations [][]int) int { //, side string, id int, parentId string) int {

	//fmt.Println(operations)

	if noValues <= 0 && value > 0 {
		return 0
	}

	if value < 0 {
		return 0
	}

	if value == 0 {
		return 1
	}

	if operations[value][noValues-1] > -1 {
		//fmt.Println("Recursion Not Called")
		return operations[value][noValues-1]
	}

	left := calc(coinValues, noValues-1, value, operations)
	right := calc(coinValues, noValues, value-coinValues[noValues-1], operations)

	operations[value][noValues-1] = left + right

	return left + right
}
