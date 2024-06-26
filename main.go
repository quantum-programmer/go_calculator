package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Map for Roman to Arabic conversion
var romanToArabic = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

// Slices for Arabic to Roman conversion
var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()
	result := calculate(input)
	fmt.Println(result)
}

func calculate(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("invalid format")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]
	var a, b int
	var isRoman bool

	if isRomanNumber(aStr) && isRomanNumber(bStr) {
		var err error
		a, err = romanToArabicConvert(aStr)
		if err != nil {
			panic(err)
		}
		b, err = romanToArabicConvert(bStr)
		if err != nil {
			panic(err)
		}
		isRoman = true
	} else if isArabicNumber(aStr) && isArabicNumber(bStr) {
		a, _ = strconv.Atoi(aStr)
		b, _ = strconv.Atoi(bStr)
		isRoman = false
	} else {
		panic("mixed number systems")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("numbers must be between 1 and 10")
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		result = a / b
	default:
		panic("invalid operator")
	}

	if isRoman {
		if result < 1 {
			panic("result is less than I in Roman numerals")
		}
		return arabicToRomanConvert(result)
	} else {
		return strconv.Itoa(result)
	}
}

func isRomanNumber(s string) bool {
	_, err := romanToArabicConvert(s)
	return err == nil
}

func isArabicNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func romanToArabicConvert(s string) (int, error) {
	total := 0
	prev := 0
	for i := 0; i < len(s); i++ {
		value, ok := romanToArabic[string(s[i])]
		if !ok {
			return 0, fmt.Errorf("invalid Roman numeral")
		}
		if value > prev {
			total += value - 2*prev
		} else {
			total += value
		}
		prev = value
	}
	return total, nil
}

func arabicToRomanConvert(num int) string {
	var result strings.Builder
	for _, pair := range arabicToRoman {
		for num >= pair.Value {
			result.WriteString(pair.Symbol)
			num -= pair.Value
		}
	}
	return result.String()
}
