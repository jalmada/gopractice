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
	fmt.Printf("%d - %d\n", value, noValues)
	fmt.Println(coinValues)
}
