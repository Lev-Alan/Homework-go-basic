package main

import (
	"fmt"
	"strings"
)
 

func main() {
	fmt.Print("_Конвертер валют_\n\n")
	for{
	currency1, number, currency2 := userInput()                         
	result := calculateResult(currency1, currency2, number)
	fmt.Printf("Ваш результат: %.02f\n\n", result)
	answer := question()
	if !answer{
		break
	}
}
}
//
//функция ввода данных пользователем
func userInput()(currency1 string, number float64, currency2 string ){              
	// ввод исходной валюты и обработка неправильного ввода
	for{
	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scanln(&currency1)
	currency1 = strings.ToUpper(currency1) // Приводим к верхнему регистру
	if currency1 != "USD"  && currency1 != "EUR" && currency1 != "RUB"{
		fmt.Printf("Вы задали неправильные параметры валюты.Повторите ввод \n")
		continue
	}
	break
}
	// ввод количества денег, которое необходимо конвертировать; обработка ошибки ввода
	for{
	fmt.Print("Введите количество денег, которое необходимо конвертировать: ")
		_, err := fmt.Scanln(&number)
        if err != nil {
            fmt.Println("Ошибка ввода! Нужно ввести число.")
            // Очищаем буфер ввода (на случай лишней строки)
            var trash string
            fmt.Scanln(&trash)
            continue
        }
	if number <=0 {
		fmt.Printf("Вы задали неправильный параметр. Повторите ввод \n")
		continue
		}
	break
}
// Вывод варианта валюты на основе предыдущего ввода
switch {
case currency1 == "USD":
	fmt.Print("Введите валюту, в которую вам необходимо конвертировать деньги (EUR, RUB): ");
case currency1 == "EUR":
	fmt.Print("Введите валюту, в которую вам необходимо конвертировать деньги (USD, RUB): ");
	default: 
	fmt.Print("Введите валюту, в которую вам необходимо конвертировать деньги (USD, EUR): ")
}
// ввод целевой валюты и обработка неправильного ввода
for{
fmt.Scanln(&currency2)
currency2 = strings.ToUpper(currency2) //Приводим к верхнему регистру

if currency2 != "EUR" && currency2!="RUB" && currency2!= "USD"{
	fmt.Printf("Вам необходимо правильно указать целевую валюту \n")
	continue
}
if currency2 == currency1{
	fmt.Println("Вам необходимо ввести целевую валюту, которая будет отличаться от исходной. \n")
	continue
}
break
}
	return currency1, number, currency2
}

// функция конвертации
const (
	USDToRUB = 81.14
	USDToEUR = 0.88
	EURToRUB = USDToRUB / USDToEUR // Рассчитанный курс EUR к RUB
)
func calculateResult(currency1, currency2 string, number float64) float64 {
	if currency1 == currency2 {
		return number
	}
	switch currency1 + "->" + currency2 {
	case "USD->EUR":
		return number * USDToEUR
	case "USD->RUB":
		return number * USDToRUB
	case "EUR->USD":
		return number / USDToEUR
	case "EUR->RUB":
		return number * EURToRUB
	case "RUB->USD":
		return number / USDToRUB
	case "RUB->EUR":
		return number / EURToRUB
	default:
		return 0
	}
}

//Вопрос о продолжении конвертации
 func question ()bool{
	fmt.Print("Хотите ли вы продолжить? (y/n): ")
	var answer string
	fmt.Scan(&answer)
	if answer == "y" || answer == "Y"{
		var trash string
        fmt.Scanln(&trash)
		return true
		
	}
	return false
 }