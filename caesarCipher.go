package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

func encryptOrDecrypt(cipherString string, initialShift int, encrypt bool) string {

	if initialShift > 25 {
		initialShift = initialShift % 26

	}

	if initialShift <= 25 {
		initialShift++
	} else {
		initialShift = 1
	}

	result := ""

	for _, v := range cipherString {

		var initialShiftedRune rune

		if encrypt {
			initialShiftedRune = v + rune(initialShift)
		} else {
			initialShiftedRune = v - rune(initialShift)
		}

		if unicode.IsLower(v) {
			if unicode.IsLower(initialShiftedRune) {
				result += string(initialShiftedRune)
			} else {
				if encrypt {
					result += string(initialShiftedRune - 26)
				} else {
					result += string(initialShiftedRune + 26)
				}
			}

		} else if unicode.IsUpper(v) {
			if unicode.IsUpper(initialShiftedRune) {
				result += string(initialShiftedRune)
			} else {
				if encrypt {
					result += string(initialShiftedRune - 26)
				} else {
					result += string(initialShiftedRune + 26)
				}
			}

		} else {
			result += string(v)
		}
	}
	return result
}

func main() {

	inputString := os.Args

	initialShiftDirecton := flag.Bool("d", false, "Flag for decryption of the passed string")
	initialinitialShiftValue := flag.Int("s", 1, "Initial shift value")
	flag.Parse()

	if *initialShiftDirecton {
		fmt.Println(encryptOrDecrypt(inputString[len(inputString)-1], *initialinitialShiftValue, false))
	} else {
		fmt.Println(encryptOrDecrypt(inputString[len(inputString)-1], *initialinitialShiftValue, true))
	}
}
