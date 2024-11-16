package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
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

	expectedTypeText := `42: int
42: int
42: int
3.14: float64
Golang: string
true: bool
(1+2i): complex64`
	expectedJoin := "4242423.14Golangtrue(1+2i)"
	expectedRune := []rune(expectedJoin)
	expectedPasswor := `4242423.14Golgo-2024angtrue(1+2i)`
	expectedHash := `53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3`

	join, rune := JoinString(s)
	password, hash := hasher(expectedRune, sault)

	assert.Equal(t, expectedTypeText, PrintType(s))
	assert.Equal(t, expectedJoin, join)
	assert.Equal(t, expectedRune, rune)
	assert.Equal(t, expectedPasswor, password)
	assert.Equal(t, expectedHash, hash)
}
