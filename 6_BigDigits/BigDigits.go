package main

import (
	"fmt"
	"log"
)

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"777777", "     7", "    7 ", "   7  ", "  7   ", " 7    ", "7     "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 999 ", "9   9", "9   9", " 9999", "    9", "    9", " 999 "},
}

func BigDigits(num string) string {
	stringOfDigits := num
	line := ""
	for row := range bigDigits[0] {
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "   "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		line += "\n"
	}
	return line
}

func main() {
	fmt.Printf("\nInput numbers: 71049285364" + "\n\n" + BigDigits("71049285364") + "\n\n")
}
