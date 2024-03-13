package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите простое математическое выражение, которое хотите вычислить. Оно должно состоять из двух целых чисел не меньше 1 и не больше 10:")
		scanner.Scan()
		input := scanner.Text()
		if scanner.Err() != nil {
			fmt.Println("Ошибка: ", scanner.Err())
		}
		// Убираем пробелы
		input = strings.ReplaceAll(input, " ", "")
		var result int
		// firstinput понадобится нам позднее
		firstinput := input
		// Основная логика
		if len(input) != 0 {
			switch {
			case strings.Contains(input, "+"):
				a := strings.Split(input, "+")
				if len(a) > 2 {
					panic("Калькулятор работает только с двумя числами")
				}
				a[0], a[1] = romanCheckAndReplace(a)
				arg1, err := strconv.Atoi(a[0])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				arg2, err := strconv.Atoi(a[1])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				if arg1 > 10 || arg2 > 10 || arg1 < 1 || arg2 < 1 {
					panic("Калькулятор не работает с числами меньше 1 или больше 10.")
				}
				result = arg1 + arg2
			case strings.Contains(input, "-"):
				a := strings.Split(input, "-")
				if len(a) > 2 {
					panic("Калькулятор работает только с двумя числами")
				}
				a[0], a[1] = romanCheckAndReplace(a)
				arg1, err := strconv.Atoi(a[0])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				arg2, err := strconv.Atoi(a[1])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				if arg1 > 10 || arg2 > 10 || arg1 < 1 || arg2 < 1 {
					panic("Калькулятор не работает с числами меньше 1 или больше 10.")
				}
				result = arg1 - arg2
			case strings.Contains(input, "*"):
				a := strings.Split(input, "*")
				if len(a) > 2 {
					panic("Калькулятор работает только с двумя числами")
				}
				a[0], a[1] = romanCheckAndReplace(a)
				arg1, err := strconv.Atoi(a[0])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				arg2, err := strconv.Atoi(a[1])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				if arg1 > 10 || arg2 > 10 || arg1 < 1 || arg2 < 1 {
					panic("Калькулятор не работает с числами меньше 1 или больше 10.")
				}
				result = arg1 * arg2
			case strings.Contains(input, "/"):
				a := strings.Split(input, "/")
				if len(a) > 2 {
					panic("Калькулятор работает только с двумя числами")
				}
				a[0], a[1] = romanCheckAndReplace(a)
				arg1, err := strconv.Atoi(a[0])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				arg2, err := strconv.Atoi(a[1])
				if err != nil {
					panic("Синтаксическая ошибка")
				}
				if arg1 > 10 || arg2 > 10 || arg1 < 1 || arg2 < 1 {
					panic("Калькулятор не работает с числами меньше 1 или больше 10.")
				}
				result = arg1 / arg2
			default:
				panic("Синтаксическая ошибка.")
			}
		} else {
			panic("Введена пустая строка")

		}
		// Переводим ответ в римские если в выражении были римские
		if strings.Contains(firstinput, "I") ||
			strings.Contains(firstinput, "V") ||
			strings.Contains(firstinput, "X") ||
			strings.Contains(firstinput, "L") ||
			strings.Contains(firstinput, "C") ||
			strings.Contains(firstinput, "D") ||
			strings.Contains(firstinput, "M") {
			if result < 1 {
				panic("В системе римских цифр нет отрицательных чисел или нуля")
			}
			fmt.Println("Ответ:", integerToRoman(result))
			break
		}
		fmt.Println("Ответ:", result)
	}

}

func romanCheckAndReplace(a []string) (string, string) {
	if (strings.Contains(a[0], "I") ||
		strings.Contains(a[0], "V") ||
		strings.Contains(a[0], "X") ||
		strings.Contains(a[0], "L") ||
		strings.Contains(a[0], "C") ||
		strings.Contains(a[0], "D") ||
		strings.Contains(a[0], "M")) &&
		(strings.Contains(a[1], "I") ||
			strings.Contains(a[1], "V") ||
			strings.Contains(a[1], "X") ||
			strings.Contains(a[1], "L") ||
			strings.Contains(a[1], "C") ||
			strings.Contains(a[1], "D") ||
			strings.Contains(a[1], "M")) {
		rim := map[string]int{
			"I": 1,
			"V": 5,
			"X": 10,
			"L": 50,
			"C": 100,
			"D": 500,
			"M": 1000,
		}

		sum := 0
		for i, v := range a[0] {
			s := a[0]
			sum += rim[string(v)]
			if i != 0 {
				if rim[string(s[i-1])] < rim[string(v)] {
					sum -= 2 * rim[string(s[i-1])]
				}
			}
		}
		a[0] = strconv.Itoa(sum)

		sum = 0
		for i, v := range a[1] {
			s1 := a[1]
			sum += rim[string(v)]
			if i != 0 {
				if rim[string(s1[i-1])] < rim[string(v)] {
					sum -= 2 * rim[string(s1[i-1])]
				}
			}
		}
		a[1] = strconv.Itoa(sum)

	} else {
		if strings.Contains(a[0], "I") ||
			strings.Contains(a[0], "V") ||
			strings.Contains(a[0], "X") ||
			strings.Contains(a[0], "L") ||
			strings.Contains(a[0], "C") ||
			strings.Contains(a[0], "D") ||
			strings.Contains(a[0], "M") ||
			strings.Contains(a[1], "I") ||
			strings.Contains(a[1], "V") ||
			strings.Contains(a[1], "X") ||
			strings.Contains(a[1], "L") ||
			strings.Contains(a[1], "C") ||
			strings.Contains(a[1], "D") ||
			strings.Contains(a[1], "M") {
			panic("Оба оператора должны быть записаны арабскими либо римскими числами")
		}
	}
	return a[0], a[1]
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
