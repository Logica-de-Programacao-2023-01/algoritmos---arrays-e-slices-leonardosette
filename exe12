package main

import (
	"fmt"
)

// Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha
// apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
func main() {
	var x int
	num := []int{}
	mul := []int{}
	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o numero %d:", i+1)
		fmt.Scan(&x)
		num = append(num, x)
	}
	for j := 0; j < len(num); j++ {
		if num[j]%3 == 0 {
			mul = append(mul, num[j])
		}
	}
	fmt.Println(num)
	fmt.Print(mul)
}
