package a1notation

import (
	"fmt"
	"math"
	"strconv"
)

var (
	numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	letters = map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z"}
)

// row and col start index 0
func FromIndex(row, col int) string {
	return fmt.Sprintf("%s%d", ColumnLetterFrom(col), row+1)
}

// col starts index 0
func ColumnLetterFrom(col int) string {
	if col >= len(letters) {
		head := ColumnLetterFrom(int(col/len(letters)) - 1)
		tail := ColumnLetterFrom(col % len(letters))
		return head + tail
	} else {
		return string(letters[col])
	}
}

func ToIndex(alpha string) (int, int) {
	numberStarts := 0
LOOP:
	for i, a := range alpha {
		a := string(a)
		for _, v := range numbers {
			if a == v {
				numberStarts = i
				break LOOP
			}
		}
	}
	row, _ := strconv.Atoi(string(alpha[numberStarts:]))
	row = row - 1

	col := 0
	i := 1
	// starting from a lower digit
	for n := len(alpha[:numberStarts]) - 1; n >= 0; n-- {
		r := string(alpha[n])
		for m, item := range letters {
			if item == r {
				col += (m + 1) * int(math.Pow(float64(len(letters)), float64(i-1)))
				break
			}
		}
		i++
	}
	col = col - 1
	return row, col

}
