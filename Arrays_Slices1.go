package main

import "fmt"

func main() {

	numeros := [3]int{2, 4, 6}
	soma := 0
	for _, valor := range numeros {
		soma += valor
	}
	fmt.Println("A soma dos valores no array é", soma)
}
