package main

import (
	"fmt"
	"unicode"
)

func MovingShift(s string, shift int) []string {

	encrypted := []string{}

	for _, v := range s {

		if unicode.IsLower(v) {
			if unicode.IsLower(v + rune(shift)) {
				encrypted = append(encrypted, string(v+rune(shift)))
			} else {
				encrypted = append(encrypted, string(v+rune(shift-26)))
			}

		} else if unicode.IsUpper(v) {
			if unicode.IsUpper(v + rune(shift)) {
				encrypted = append(encrypted, string(v+rune(shift)))
			} else {
				encrypted = append(encrypted, string(v+rune(shift-26)))
			}

		} else {
			encrypted = append(encrypted, string(v))
		}
		if shift <= 25 {
			shift++
		} else {
			shift = 1
		}
	}

	//return array of 5 coded strings
	return encrypted
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