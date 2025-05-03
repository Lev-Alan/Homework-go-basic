package main

import (
	"fmt"
)
const USDtoRUB = 81.14
const USDtoEUR = 0.88  

func main() {
	fmt.Print("_Конвертер валют_\n\n")
	for{
	currency1, number, currency2 := userInput()                         
	result := calculateResult(currency1, currency2, number)
	fmt.Printf("Ваш результат: %.04f\n\n", result)
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
	if currency1 != "USD"  && currency1 != "EUR" && currency1 != "RUB"{
		fmt.Printf("Вы задали неправильные параметры валюты.Повторите ввод \n\n")
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
		fmt.Printf("Вы задали неправильный параметр. Повторите ввод \n\n")
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
if currency2 != "EUR" && currency2!="RUB" && currency2!= "USD"{
	fmt.Printf("Вам необходимо правильно указать целевую валюту \n\n")
	continue
}
break
}
	return currency1, number, currency2
}

// функция конвертации
func calculateResult(currency1, currency2 string, number float64)(result float64){
	switch{
	case currency1 == "USD" && currency2 == "EUR":
		result = USDtoEUR * number;
	case currency1 == "USD" && currency2 == "RUB":
		result = USDtoRUB * number;
	case currency1 == "EUR" && currency2 == "USD":
		result = number/USDtoEUR;
	case currency1 == "EUR" && currency2 == "RUB":
		result = number * USDtoRUB / USDtoEUR;
	case currency1 == "RUB" && currency2 == "USD":
		result = number / USDtoRUB;
	case currency1 == "RUB" && currency2 == "EUR":
		result = USDtoEUR * number / USDtoRUB
	}
	return result
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
