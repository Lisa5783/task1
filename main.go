package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите выражение: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if input == "0" {
			fmt.Println("Программа завершена")
			return
		} else {
			elements := strings.Fields(input)
			if len(elements) != 3 {
				fmt.Println("Неверный формат ввода")
				continue
			}

			num1, err1 := strconv.ParseFloat(elements[0], 64)
			num2, err2 := strconv.ParseFloat(elements[2], 64)

			if err1 != nil || err2 != nil {
				fmt.Println("Ошибка при чтении чисел")
				continue
			}

			var result float64
			switch elements[1] {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				if num2 == 0 {
					fmt.Println("Деление на ноль")
					continue
				}
				result = num1 / num2
			default:
				fmt.Println("Неподдерживаемая операция")
				continue
			}

			fmt.Printf("%s %s %s = %.2f\n", elements[0], elements[1], elements[2], result)
		}
	}
}
