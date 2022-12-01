package main

import (
	"fmt"
	"unicode/utf8"
)

/*
A Go string is a read-only slice of bytes.
In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
*/

func main() {

	const s = "世界你好" // 一个中文占三个字符

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // hex value 十六进制
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}
func examineRune(r rune) {

	if r == 't' {
		fmt.Println("find t")
	} else if r == 'ส' {
		fmt.Println("find sp sua")
	}

}
