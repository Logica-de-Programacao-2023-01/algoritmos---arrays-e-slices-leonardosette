package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice Original", slice)

	indice := 2 // indice terceiro elemento
	slice = append(slice[:indice], slice[indice+1:]...)
	fmt.Println("Slice após a remoção do terceiro elemento", slice)

}
