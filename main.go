package main

import (
	"fmt"
	"unicode"
)

func encryptOrDecrypt(s string, shift int, encrypt bool) string {

	result := ""

	for _, v := range s {

		var shiftedRune rune

		if encrypt {
			shiftedRune = v + rune(shift)
		} else {
			shiftedRune = v - rune(shift)
		}

		if unicode.IsLower(v) {
			if unicode.IsLower(shiftedRune) {
				result += string(shiftedRune)
			} else {
				if encrypt {
					result += string(shiftedRune - 26)
				} else {
					result += string(shiftedRune + 26)
				}
			}

		} else if unicode.IsUpper(v) {
			if unicode.IsUpper(shiftedRune) {
				result += string(shiftedRune)
			} else {
				if encrypt {
					result += string(shiftedRune - 26)
				} else {
					result += string(shiftedRune + 26)
				}
			}

		} else {
			result += string(v)
		}
		if shift <= 25 {
			shift++
		} else {
			shift = 1
		}
	}
	return result
}

func main() {

	inputString := os.Args

	shiftDirecton := flag.Bool("d", false, "used to decrypt the passed string")
	flag.Parse()

	if *shiftDirecton {
		fmt.Println(encryptOrDecrypt(inputString[2], 1, false))
	} else {
		fmt.Println(encryptOrDecrypt(inputString[1], 1, true))
	}

}