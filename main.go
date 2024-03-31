package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	f "employee/Func"
)


func main() {
	signs := []string{"+", "-", "*", "/"}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		contains, characters := f.CheckString(text, signs)

		if contains {
			data := strings.Split(text, characters)

			if len(data) != 2 {
				log.Panic("Не корректный ввод")
			}
			
			fmt.Println("", f.FindMathMethod(data, characters))
		} else {
			fmt.Println("Не правильный формат!")
		}
	}
}
