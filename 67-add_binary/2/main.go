package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addBinary("0", "1"))
}

func addBinary(a string, b string) string {
	var s strings.Builder
	// carry := 0
	lenA := len(a)
	lenB := len(b)

	for lenA > 0 && lenB > 0 {
		temp := int(a[lenA-1]-'0') + int(b[lenB-1]-'0')
		fmt.Println("temp: ", temp)
		s.WriteByte(byte(temp + '0'))
		lenA--
		lenB--
	}

	return s.String()
}
