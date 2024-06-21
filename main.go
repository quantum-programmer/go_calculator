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
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Map for Arabic to Roman conversion
var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	result, err := calculate(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func calculate(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid format")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]
	var a, b int
	var isRoman bool

	if isRomanNumber(aStr) && isRomanNumber(bStr) {
		a = romanToArabic[aStr]
		b = romanToArabic[bStr]
		isRoman = true
	} else if isArabicNumber(aStr) && isArabicNumber(bStr) {
		a, _ = strconv.Atoi(aStr)
		b, _ = strconv.Atoi(bStr)
		isRoman = false
	} else {
		return "", fmt.Errorf("mixed number systems")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		return "", fmt.Errorf("numbers must be between 1 and 10")
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
			return "", fmt.Errorf("division by zero")
		}
		result = a / b
	default:
		return "", fmt.Errorf("invalid operator")
	}

	if isRoman {
		if result < 1 {
			return "", fmt.Errorf("result is less than I in Roman numerals")
		}
		return arabicToRoman[result], nil
	} else {
		return strconv.Itoa(result), nil
	}
}

func isRomanNumber(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

func isArabicNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
