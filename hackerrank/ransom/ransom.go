package ransom

import (
    "fmt"
    "os"
)

func Run(){
    var l1, l2 string
	fmt.Fscanln(os.Stdin, &l1)
	fmt.Fscanln(os.Stdin, &l2)

	fmt.Println(l1 + l2)
}