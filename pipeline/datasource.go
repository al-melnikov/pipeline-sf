// Здесь описан источник данных для пайплайна
package pipeline

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Считывает целые числа с клавиатуры
func ReceiveData() (<-chan int, <-chan bool) {
	res := make(chan int)
	done := make(chan bool)

	go func() {
		defer close(done)
		scanner := bufio.NewScanner(os.Stdin)
		var data string
		for {
			scanner.Scan()
			data = scanner.Text()
			if strings.EqualFold(data, "exit") || strings.EqualFold(data, "quit") {
				fmt.Println("Программа завершила работу!")
				return
			}
			i, err := strconv.Atoi(data)
			if err != nil {
				fmt.Println("Программа обрабатывает только целые числа!")
				continue
			}
			res <- i
		}
	}()
	return res, done
}
