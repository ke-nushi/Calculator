package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(num1, num2 int) int {

	return num1 + num2
}

func sub(num1, num2 int) int {

	return num1 - num2

}

func multy(num1, num2 int) int {

	return num1 * num2

}

func div(num1, num2 int) int {

	return num1 / num2

}

func findOperator(str string) string {

	switch {

	case strings.Contains(str, "+"):
		return "+" //, nil

	case strings.Contains(str, "-"):
		return "-" //, nil

	case strings.Contains(str, "*"):
		return "*" //, nil

	case strings.Contains(str, "/"):
		return "/" //, nil

	default:
		panic("строка не является математической операцией")

	}

}

func calculate(num1, num2 int, operator string) (num int) {

	switch operator {

	case "+":

		num = sum(num1, num2)

	case "-":

		num = sub(num1, num2)

	case "*":

		num = multy(num1, num2)

	case "/":

		num = div(num1, num2)

	default:
		//err = fmt.Errorf("Выдача паники, так как строка не является математической операцией", operator)
		panic("строка не является математической операцией")
	}
	return
}

var Roman = []string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func isRoman(num1 string) bool {

	for i := 0; i < len(Roman); i++ {
		if num1 == Roman[i] {
			return true
		}
	}
	return false
}
func romanToInt(nums string) int {
	i := 0
	for ; i < len(Roman); i++ {
		if nums == Roman[i] {
			return i
		}

	}
	return i
}

func intToRoman(res int) string {
	romanres := Roman[res]
	return romanres
}

func getNumsAndType(text, op string) (a, b int, Rom bool) {

	nums := strings.Split(text, op)

	if len(nums) > 2 {

		panic("Больше 2 аргументов")
	}
	firstRomType := isRoman(nums[0])
	secondRomType := isRoman(nums[1])

	if firstRomType != secondRomType {
		panic("Разные форматы чисел")

	}

	if firstRomType && secondRomType {
		Rom = true
		a = romanToInt(nums[0])
		b = romanToInt(nums[1])

	} else {
		a, _ = strconv.Atoi(nums[0])

		b, _ = strconv.Atoi(nums[1])

	}
	if a < 1 || a > 10 || b < 1 || b > 10 {

		panic("Число не может быть меньше 1 или больше 10")
	}
	return a, b, Rom
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input:")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")

		operator := findOperator(text)
		a, b, isRom := getNumsAndType(text, operator)
		result := calculate(a, b, operator)

		if isRom {
			if result <= 0 {
				panic("Римские числа не могут быть меньше нуля")
			}
			//first := intToRoman(a)
			//second := intToRoman(b)
			res := intToRoman(result)

			fmt.Println("Output:", "\n", res)
		} else {
			fmt.Println("Output:", "\n", result)
		}
	}
}
