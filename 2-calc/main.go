package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := operation()
	numbers := numberInput()
	calculateOperation(operation, numbers)
}

func operation()( string){
	fmt.Print("Выберите операцию, которую вы хотите совершить (AVG-cреднее, SUM-сумма, MED-медиана чисел): ")
	var operation string
	for{
		operation, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		operation = strings.ToUpper(strings.TrimSpace(operation))
			switch operation{
			case "AVG", "SUM", "MED":
				break
			default:
				fmt.Println("Вы ввели неверное значение. Пожалyйста, повторите ввод: ")
				continue
			}
			break
	}
	return operation
}

func numberInput() []int {
	var numbers []int
	fmt.Print("Введите набор целых чисел через запятую: ")
	reader := bufio.NewReader(os.Stdin)
	for{
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		separateInput := strings.Split(input, ",")
		numberSlice := make([]int, 0,len(separateInput))
		inputValid := true
		for _, value := range separateInput{
			value = strings.TrimSpace(value)
			number, err := strconv.Atoi(value)
			if err!=nil{
				fmt.Print("Вы ввели неправильные значения. Повторите ввод: ")
				inputValid = false
				break
			}
			numberSlice = append(numberSlice, number)
		}
		if inputValid{
			numbers = numberSlice
			break
		}
	}
	return numbers
}

func calculateOperation(operation string, numbers []int){

	switch operation{
	case "SUM": 
		sum := calculateSum(numbers)
		fmt.Print("Результат операции: ", sum)
	case "AVG": 
		avg := calculateSum(numbers)/len(numbers)
		fmt.Print("Результат операции: ", avg)
	case "MED": 
		med := calculateMED(numbers)
		fmt.Print("Результат операции: ", med)
	}

}


//Функция расчёта суммы
func calculateSum(numbers []int)(sum int){    
	sum = 0
	for _, value := range numbers{
		sum +=value
	}
	return 
}
//Функция для расчёта медианы чисел
func calculateMED(numbers []int)float64{
	n:=len(numbers)
	sort.Ints(numbers)
	if n%2==1{
		return float64(n/2)
	}
	return float64(((n/2-1)+(n/2))/2)
	
}
