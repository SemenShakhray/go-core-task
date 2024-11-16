package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func PrintType(s []any) string {
	var text string
	for i, v := range s {
		if i == len(s)-1 {
			text += fmt.Sprintf("%v: %T", v, v)
		} else {
			text += fmt.Sprintf("%v: %T\n", v, v)
		}
	}
	return text

}

func JoinString(s []any) (string, []rune) {
	var str strings.Builder
	for _, v := range s {
		str.WriteString(fmt.Sprint(v))
	}

	strJoin := str.String()
	strRune := []rune(strJoin)
	fmt.Println(len(strJoin))
	return strJoin, strRune

}

func hasher(r []rune, sault string) (string, string) {
	var password []rune
	password = append(password, r[:len(r)/2]...)
	password = append(password, []rune(sault)...)
	password = append(password, r[len(r)/2:]...)

	hasher := sha256.New()
	hasher.Write([]byte(string(password)))
	return string(password), hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	var (
		numDecimal     int       = 42       // Десятичная система
		numOctal       int       = 052      // Восьмеричная система
		numHexadecimal int       = 0x2A     // Шестнадцатиричная система
		pi             float64   = 3.14     // Тип float64
		name           string    = "Golang" // Тип string
		isActive       bool      = true     // Тип bool
		complexNum     complex64 = 1 + 2i   // Тип complex64
	)
	var s []any
	sault := "go-2024"
	s = append(s, numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println(PrintType(s))
	join, rune := JoinString(s)
	fmt.Printf("String: %s\nRune: %v\n", join, rune)
	password, hash := hasher(rune, sault)
	fmt.Printf("PASSWORD: %s\nHASH: %s\n", password, hash)
}
