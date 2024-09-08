package main

import (
 "fmt"
 "strconv"
 "strings"
)

var romanNumerals = map[rune]int{
 'I': 1,
 'V': 5,
 'X': 10,
 'L': 50,
 'C': 100,
}

func romanToArabic(roman string) (int, error) {
 total := 0
 lastValue := 0
 for _, r := range roman {
  value, exists := romanNumerals[r]
  if !exists {
   return 0, fmt.Errorf("некорректный римский символ: %c", r)
  }

  if value > lastValue {
   total += value - 2*lastValue 
  } else {
   total += value
  }
  lastValue = value
 }
 return total, nil
}

func arabicToRoman(num int) (string, error) {
 if num < 1 || num > 100 {
  return "", fmt.Errorf("число должно быть в диапазоне от 1 до 100")
 }

 var result strings.Builder
 romanPairs := []struct {
  Value int
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
 return result.String(), nil
}

func calculate(a, b int, operation string) int {
 switch operation {
 case "+":
  return a + b
 case "-":
  return a - b
 case "*":
  return a * b
 case "/":
  return a / b
 default:
  panic("недопустимая операция")
 }
}

func main() {
 var firstInput, secondInput, operation string

 fmt.Println("Введите первое число (римское или арабское): ")
 fmt.Scan(&firstInput)

 fmt.Println("Введите оператор (+, -, *, /): ")
 fmt.Scan(&operation)

 fmt.Println("Введите второе число (римское или арабское): ")
 fmt.Scan(&secondInput)

 
 firstValue, err1 := strconv.Atoi(firstInput)
 if err1 != nil {
  
  firstValue, err1 = romanToArabic(firstInput)
  if err1 != nil {
   fmt.Println(err1)
   return
  }
 }

 secondValue, err2 := strconv.Atoi(secondInput)
 if err2 != nil {
  secondValue, err2 = romanToArabic(secondInput)
  if err2 != nil {
   fmt.Println(err2)
   return
  }
 }

 result := calculate(firstValue, secondValue, operation)

 romanResult, err := arabicToRoman(result)
 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Printf("Результат: %d (%s)n", result, romanResult)
}
	



