package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const( 
	USDToRUB = 81.14
	USDtoEUR = 0.88
	EURToRUB = USDToRUB/USDtoEUR
)

func main() {
	fmt.Println("_Конвертер валют_")
	number := enterNumber()
	exchangeRates := map[string]map[string]float64{
		"RUB": {
			"EUR": number * USDtoEUR / USDToRUB,
			"USD": number / USDToRUB,
		},
		"USD": {
			"RUB": number * USDToRUB,
			"EUR": number * USDtoEUR,
		},
		"EUR": {
			"RUB": number * USDToRUB / USDtoEUR,
			"USD": number / USDtoEUR,
		},
	}
	currency1, currency2 :=currencylInput(exchangeRates)
	fmt.Printf("%0.0f %s = %0.2f %s",number,currency1,exchangeRates[currency1][currency2],currency2)
}
// Получаем данные о конвертируемой сумме денег
func enterNumber()(number float64){
	fmt.Println("Введите сумму, которую необходимо конвертировать. ")
	for{
		fmt.Scan(&number)
		fmt.Scanln(new(string))
		if number <=0 {
			fmt.Println("\nВы задали неправильные параметры. Пожалуйста, повторите ввод.")
			continue
		}
		break
	}
	return number
}

//Функция ввода валюты с валидацией
func currencylInput(exchangeRates map[string]map[string]float64 )(string, string){
	errorInput:=fmt.Sprint("Ошибка. Выберите вариант из списка: ")
	var currency1, currency2 string
	fmt.Println("Введите исходную валюту")
	//Вывод доступных вариантов исходной валюты
	for key := range exchangeRates{
		fmt.Println("-", key)
	}
	for{
		currency1=input()
		if _, ok := exchangeRates[currency1]; ok{
			break
		}
		fmt.Print(errorInput)
	}
	fmt.Println("В какую валюту конвертировать: ")
	//Вывод доступных вариантов целевой валюты
	for key := range exchangeRates[currency1]{
		fmt.Println("-", key)
	}
	for{
		currency2=input()
		if _, ok := exchangeRates[currency1][currency2]; ok{
			break
		}
		fmt.Print(errorInput)
	}
		return currency1, currency2
}

// Универсальный запрос на ввод валюты
func input()string{
	scaner:=bufio.NewScanner(os.Stdin)
	scaner.Scan()
	input:=strings.ToUpper(strings.TrimSpace(scaner.Text()))
	return input
}



