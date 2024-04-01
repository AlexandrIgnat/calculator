package Func

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

func Alert() { fmt.Println("Много математических символов") }

func CheckString(input string, array []string) (is_contains bool, symbol string) {
	is_contains = false
	for _, char := range array {
		if strings.Contains(input, char) && is_contains == false {
			is_contains = true
			symbol = char
		} else if strings.Contains(input, char) && is_contains == true {
			Alert()
			os.Exit(0)
		}
	}
	return
}

func Add(val_left, val_right int) int {
	return val_left + val_right
}

func Sub(val_left, val_right int) int {
	return val_left - val_right
}

func Multy(val_left, val_right int) int {
	return val_left * val_right
}

func Div(val_left, val_right int) int {
	return val_left / val_right
}

func validation(data []string) (val_left, val_right int) {
	some1, left_err := strconv.ParseFloat(strings.TrimSpace(data[0]), 64)
	some2, right_err := strconv.ParseFloat(strings.TrimSpace(data[1]), 64)
	val_left = int(math.Round(some1))
	val_right = int(math.Round(some2))
	if left_err != nil || right_err != nil || val_left > 10 || val_right > 10 {
		log.Panic("Не корректный ввод")
	}

	return
}

func is_rome_numbers(data []string) (val_left, val_right int, is_rome_numbers bool) {
	rom_numbers := map[string]int{
		"Ⅰ": 1, "Ⅱ": 2, "III": 3,
		"IV": 4, "V": 5, "VI": 6,
		"VII": 7, "VIII": 8, "IX": 9,
		"X": 10,
	}
	val_left, left_err := rom_numbers[strings.TrimSpace(data[0])]
	val_right, right_err := rom_numbers[strings.TrimSpace(data[1])]
		
	if (left_err && !right_err) || (!left_err && right_err) {
		log.Panic("Не корректный ввод")
	} else if !left_err && !right_err {
		is_rome_numbers = false
	} else {
		is_rome_numbers = true
	}

	return
}

func intToRoman(romanResult int) (string) {
	var romanNum string
	if romanResult == 0 {
		panic("В римской системе нет числа 0.")
	} else if romanResult < 0 {
		panic("В римской системе нет отрицательных чисел.")
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}

	return romanNum
}

func FindMathMethod(data []string, characters string) string {
	var val int
	val_left, val_right, is_rome_numbers := is_rome_numbers(data)
	if !is_rome_numbers {
		val_left, val_right := validation(data)
		switch characters {
			case "+":
				val = Add(val_left, val_right)
			case "-":
				val = Sub(val_left, val_right)
			case "*":
				val = Multy(val_left, val_right)
			case "/":
				val = Div(val_left, val_right)
			default:
				panic("Не распознанный знак.")
			}

		return strconv.Itoa(val)
	} else {
		switch characters {
			case "+":
				val = Add(val_left, val_right)
			case "-":
				val = Sub(val_left, val_right)
			case "*":
				val = Multy(val_left, val_right)
			case "/":
				val = Div(val_left, val_right)
			default:
				panic("Не распознанный знак.")
		}

		return intToRoman(val)
	}
}
