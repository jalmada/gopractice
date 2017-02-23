package median

import (
	"fmt"

	"github.com/xaratustra/utils"
)

//Run Median test
func Run() {
	heap := utils.NewMinHeap()
	heap.Print()
	fmt.Printf("Add: %v\n", 8)
	heap.Add(8)
	heap.Print()
	val, _ := heap.Poll()
	fmt.Printf("Removed: %v\n", val)
	heap.Print()
	fmt.Printf("Add: %v\n", 25)
	heap.Add(25)
	heap.Print()
	val, _ = heap.Poll()
	fmt.Printf("Removed: %v\n", val)
	heap.Print()

}
