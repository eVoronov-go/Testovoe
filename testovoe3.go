package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, bool) {

	var romanNumerals = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10}

	isRoman := false

	resNum, _ := strconv.Atoi(roman)

	if resNum == 0 {
		resNum = romanNumerals[roman]
		isRoman = true
		return resNum, isRoman
	}

	if resNum < 1 || resNum > 10 {
		panic("Число должно быть в диапазоне от 1 до 10!")
	}
	return resNum, isRoman
}

func arabicToRoman(num int) string {
	var result strings.Builder
	romanPairs := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, pair := range romanPairs {
		for num >= pair.Value {
			result.WriteString(pair.Symbol)
			num -= pair.Value
		}
	}
	return result.String()
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2

	case "-":
		return num1 - num2

	case "*":
		return num1 * num2

	case "/":
		if num2 == 0 {
			panic("Деление на ноль!")
		}
		return num1 / num2
	default:
		panic("Неверная операция")
	}
}

func resultCalc(Text string) string {
	operators := []string{"+", "-", "*", "/"}
	Text = strings.Trim(Text, "\r\n")

	if Text == "Выйти" {
		os.Exit(0)
	}

	Text = strings.ReplaceAll(Text, " ", "")

	var operator string

	for _, op := range operators {
		if strings.Contains(Text, op) {
			operator = op
			if strings.Contains(Text, op) == false {
				panic("Оператор не найден")
			}
			break
		}
	}

	nums := strings.Split(Text, operator)

	if len(nums) > 2 {
		panic("Большое количество операндов!")
	}

	var numRoman1, numRoman2 bool
	num1, numRoman1 := romanToArabic(nums[0])
	num2, numRoman2 := romanToArabic(nums[1])

	if num1 == 0 || num2 == 0 {
		panic("Операнд не в диапазоне счисления")
	}

	var result int

	if numRoman1 == false && numRoman2 == true {
		panic("Разные системы счисления!")
	} else if numRoman1 == true && numRoman2 == false {
		panic("Разные системы счисления!")
	} else if numRoman1 == true && numRoman2 == true {
		result = calculate(num1, num2, operator)
		if result <= 0 {
			panic("Неверное римское число")
		}
		return arabicToRoman(result)
	}

	result = calculate(num1, num2, operator)

	resultRoman := strconv.Itoa(result)
	return resultRoman
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Если хотите выйти, введите: Выйти ")
		fmt.Println("Введите математическую операцию:")
		inputText, _ := scanner.ReadString('\n')

		fmt.Println("Результат вычисления: ", resultCalc(inputText))
	}
}
