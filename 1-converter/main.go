package main

import "fmt"

const USDtoRUB = 81.14
const USDtoEUR = 0.88          
func main() {
	fmt.Println(" Конвертер валют ")
	n := getUSerInput() 
	EUR := EURconverter(n, USDtoRUB, USDtoEUR)   // количество евро, которое нужно конвертировать в рубли
	outputResult(EUR)
}
func getUSerInput ()(n float64){
	fmt.Println("Введите количество евро, которое необходимо конвертировать в рубли")
	fmt.Scan(&n)
	return n
}
func EURconverter (n, USDtoRUB, USDtoEUR float64)(float64){
	EUR := n * USDtoRUB / USDtoEUR // формула для конвертирования евро рубли
	return EUR
	
}
func outputResult(EUR float64){
	result := fmt.Sprintf("Ваш результат в рублях: %.02f", EUR)
	fmt.Print(result)
}