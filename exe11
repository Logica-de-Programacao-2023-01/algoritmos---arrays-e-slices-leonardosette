package main

import "fmt"

func main() {
	var linha, coluna int
	matriz := [2][3]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Elemento[%d][%d] = ", i+1, j+1)
			fmt.Scan(&matriz[i][j])
		}
	}
	fmt.Print("Digite o numero da linha:")
	fmt.Scan(&linha)
	fmt.Print("Digite o numero da coluna:")
	fmt.Scan(&coluna)
	fmt.Print(matriz[linha-1][coluna-1])
}
