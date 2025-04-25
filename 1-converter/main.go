package main

import "fmt"

func main() {
	const USDtoEUR = 0.88                   // константа конвертации из 1 USD to EUR
	const USDtoRUB = 81.14                  // константа конвертации из 1 USD to RUB
	fmt.Println("Конвертер валют")
	n := 3.0                                // n количество евро, которое нам необходимо конвертировать в рубли
	EUR := n * USDtoRUB / USDtoEUR          // формула для конвертации Евро в рубли
	fmt.Printf("Ваш результат в рублях: %.02f", EUR)

}