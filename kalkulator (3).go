package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rimskie_v_arabskie = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabskie_v_rimskie = []struct {
	value int
	roman string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"},
	{9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Преобразование римских чисел в арабские
func convert_rimskie_v_arabskie(rimskie string) (int, error) {
	value, ok := rimskie_v_arabskie[strings.ToUpper(rimskie)]
	if !ok {
		return 0, errors.New("неизвестное римское число")
	}
	return value, nil
}

// Преобразование арабских чисел в римские
func convert_arabskie_v_rimskie(arabskie int) (string, error) {
	if arabskie < 1 {
		return "", errors.New("результат римского числа должен быть больше или равен 1")
	}
	var result string
	for _, entry := range arabskie_v_rimskie {
		for arabskie >= entry.value {
			result += entry.roman
			arabskie -= entry.value
		}
	}
	return result, nil
}

// Проверка на диапазон чисел
func validateNumber(num int) error {
	if num < 1 || num > 10 {
		return errors.New("число должно быть в диапазоне от 1 до 10")
	}
	return nil
}

// Арифметические операции
func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на 0 невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}

// Проверка на римские числа
func isRoman(input string) bool {
	_, err := convert_rimskie_v_arabskie(input)
	return err == nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, err := reader.ReadString('\n')
	if err != nil {
		panic("Ошибка: неверный формат ввода.")
	}

	// Удаление лишних пробелов и символов новой строки
	expression = strings.TrimSpace(expression)

	// Разбиваем строку на операнды и оператор
	parts := strings.Split(expression, " ")

	// Проверяем, что введено ровно 3 части (два числа и один оператор)
	if len(parts) != 3 {
		panic("Ошибка: неверный формат выражения. Должно быть два операнда и один оператор.")
	}

	firstInput, operator, secondInput := parts[0], parts[1], parts[2]

	isFirstRoman := isRoman(firstInput)
	isSecondRoman := isRoman(secondInput)

	// Проверка на смешение римских и арабских чисел
	if isFirstRoman != isSecondRoman {
		panic("Ошибка: нельзя смешивать римские и арабские числа")
	}

	var a, b int
	if isFirstRoman {
		// Преобразуем римские числа в арабские
		a, err = convert_rimskie_v_arabskie(firstInput)
		if err != nil {
			panic(err)
		}
		b, err = convert_rimskie_v_arabskie(secondInput)
		if err != nil {
			panic(err)
		}
	} else {
		// Преобразуем арабские числа
		a, err = strconv.Atoi(firstInput)
		if err != nil {
			panic("Ошибка: неверное первое число.")
		}
		b, err = strconv.Atoi(secondInput)
		if err != nil {
			panic("Ошибка: неверное второе число.")
		}
	}

	// Валидация диапазона чисел
	if err := validateNumber(a); err != nil {
		panic(err)
	}
	if err := validateNumber(b); err != nil {
		panic(err)
	}

	// Вычисление результата
	result, err := calculate(a, b, operator)
	if err != nil {
		panic(err)
	}

	// Вывод результата
	if isFirstRoman {
		// Если числа римские, результат должен быть больше 0
		if result < 1 {
			panic("Ошибка: результат римского числа должен быть больше 0")
		}
		romanResult, err := convert_arabskie_v_rimskie(result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Результат: %s\n", romanResult)
	} else {
		// Выводим результат для арабских чисел
		fmt.Printf("Результат: %d\n", result)
	}
}
