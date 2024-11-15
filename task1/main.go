package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Printf("numDecimal: %T\n", numDecimal)
	fmt.Printf("numOctal: %T\n", numOctal)
	fmt.Printf("numHexadecimal: %T\n", numHexadecimal)
	fmt.Printf("pi: %T\n", pi)
	fmt.Printf("name: %T\n", name)
	fmt.Printf("isActive: %T\n", isActive)
	fmt.Printf("complexNum: %T\n", complexNum)

	var str strings.Builder
	str.WriteString(fmt.Sprint(numDecimal))
	str.WriteString(fmt.Sprint(numOctal))
	str.WriteString(fmt.Sprint(numHexadecimal))
	str.WriteString(fmt.Sprint(pi))
	str.WriteString(fmt.Sprint(name))
	str.WriteString(fmt.Sprint(isActive))
	str.WriteString(fmt.Sprint(complexNum))

	strJoin := str.String()
	strRune := []rune(strJoin)
	fmt.Printf("Result string: %s\nRune: %v\n", strJoin, strRune)

	sault := "go-2024"
	var password []rune
	password = append(password, strRune[:len(strRune)/2]...)
	password = append(password, []rune(sault)...)
	password = append(password, strRune[len(strRune)/2:]...)

	hasher := sha256.New()
	hasher.Write([]byte(string(password)))
	fmt.Printf("PASSWORD: %s\nHASH: %s\n", string(password), hex.EncodeToString(hasher.Sum(nil)))

}
