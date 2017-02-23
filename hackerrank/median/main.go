package median

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/xaratustra/utils"
)

//Run Median test
func Run() {
	leftHeap := utils.NewMaxHeap()
	rightHeap := utils.NewMinHeap()

	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	if num == 0 {
		return
	}
	for scanner.Scan() {

		val, _ := strconv.Atoi(scanner.Text())
		if _, isEmpty := leftHeap.Peek(); isEmpty {
			leftHeap.Add(val)
		} else if leftRoot, _ := leftHeap.Peek(); val < leftRoot {
			if leftHeap.Size() > rightHeap.Size() {
				valToMove, _ := leftHeap.Poll()
				rightHeap.Add(valToMove)
			}

			leftHeap.Add(val)
		} else {
			if leftHeap.Size() < rightHeap.Size() {
				valToMove, _ := rightHeap.Poll()
				leftHeap.Add(valToMove)
			}

			rightHeap.Add(val)
		}

		leftPeek, _ := leftHeap.Peek()
		rightPeek, _ := rightHeap.Peek()
		if leftHeap.Size() == rightHeap.Size() {
			fmt.Printf("%.1f\n", float32(leftPeek+rightPeek)/2)
		} else if leftHeap.Size() < rightHeap.Size() {
			fmt.Printf("%.1f\n", float32(rightPeek))
		} else {
			fmt.Printf("%.1f\n", float32(leftPeek))
		}
		count = count + 1
		if count == num {
			break
		}
	}
}
