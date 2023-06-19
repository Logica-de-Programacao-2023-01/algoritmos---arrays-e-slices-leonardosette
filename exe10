package main

import (
	"fmt"
	"sort"
)

// Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.
func main() {
	var num int
	slice := []int{}
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o numero %d:", i+1)
		fmt.Scan(&num)
		slice = append(slice, num)
	}
	sort.Ints(slice)
	fmt.Printf("%d,%d", slice[0], slice[9])
}
