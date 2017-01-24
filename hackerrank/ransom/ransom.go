package ransom

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	bio := bufio.NewReader(os.Stdin)

	var m, n int

	fmt.Fscanf(bio, "%d %d\n", &m, &n)

	line, _ := bio.ReadString('\n')
	line2, _ := bio.ReadString('\n')

	line = strings.Replace(line, "\n", "", -1)
	line2 = strings.Replace(line2, "\n", "", -1)

	fmt.Println(line)
	fmt.Println(line2)

	//fmt.Println(m + n)

}
