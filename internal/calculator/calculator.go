package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Evaluate принимает арифметическое выражение и возвращает результат или ошибку
func Evaluate(expression string) (float64, error) {
	// Удаляем пробелы
	expression = strings.ReplaceAll(expression, " ", "")
	
	// Проверяем на пустую строку
	if expression == "" {
		return 0, errors.New("empty expression")
	}

	// Проверяем допустимые символы и формат
	for i, char := range expression {
		if !(char >= '0' && char <= '9') && 
		   !strings.ContainsRune("+-", char) &&
		   !(char == '-' && (i == 0 || expression[i-1] == '+' || expression[i-1] == '-')) {
			return 0, fmt.Errorf("invalid character in expression: %c", char)
		}
		// Проверяем на последовательные операторы
		if i > 0 && isOperator(char) && isOperator(rune(expression[i-1])) {
			return 0, errors.New("invalid format: consecutive operators")
		}
	}

	// Проверяем, не заканчивается ли выражение на оператор
	if len(expression) > 0 && isOperator(rune(expression[len(expression)-1])) {
		return 0, errors.New("invalid format: expression ends with operator")
	}

	// Обрабатываем отрицательные числа в начале
	if expression[0] == '-' {
		expression = "0" + expression
	}

	// Разбиваем по операторам, сохраняя их
	var numbers []string
	var operators []string
	current := ""
	
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || (expression[i] == '-' && i > 0) {
			if current != "" {
				numbers = append(numbers, current)
				current = ""
			}
			operators = append(operators, string(expression[i]))
		} else {
			current += string(expression[i])
		}
	}
	if current != "" {
		numbers = append(numbers, current)
	}

	// Проверяем, что у нас есть хотя бы одно число
	if len(numbers) == 0 {
		return 0, errors.New("invalid format: no numbers found")
	}

	// Вычисляем результат
	result, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", numbers[0])
	}

	for i := 0; i < len(operators); i++ {
		if i+1 >= len(numbers) {
			return 0, errors.New("invalid format: missing number after operator")
		}
		
		num, err := strconv.ParseFloat(numbers[i+1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", numbers[i+1])
		}

		switch operators[i] {
		case "+":
			result += num
		case "-":
			result -= num
		}
	}

	return result, nil
}

// isOperator проверяет, является ли символ оператором
func isOperator(c rune) bool {
	return c == '+' || c == '-'
}
