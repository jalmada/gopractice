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
	result := "No"
	fmt.Fscanf(bio, "%d %d\n", &m, &n)

	magazine, _ := bio.ReadString('\n')
	note, _ := bio.ReadString('\n')

	magWords := strings.Split(strings.Replace(strings.Trim(magazine, " "), "\n", "", -1), " ")
	noteWords := strings.Split(strings.Replace(strings.Trim(note, " "), "\n", "", -1), " ")

	if len(noteWords) > len(magWords) {
		fmt.Println(result)
		return
	}

	noteMap := make(map[string]int)
	for _, nw := range noteWords {
		if nw != "" {
			if val, ok := noteMap[nw]; ok {
				noteMap[nw] = val + 1
			} else {
				noteMap[nw] = 1
			}
		}
	}

	countNoteWords := len(noteWords)
	for _, mw := range magWords {
		if val, ok := noteMap[mw]; ok && val > 0 {
			noteMap[mw] = val - 1
			countNoteWords = countNoteWords - 1

			if countNoteWords == 0 {
				result = "Yes"
				break
			}
		}

	}

	fmt.Println(result)
}
