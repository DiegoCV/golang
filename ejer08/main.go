package main

import "fmt"

var tabla [5]int8
var matriz [5][7]int8

func main() {
	tabla[3] = 9
	for _, num := range tabla {
		fmt.Println(num)
	}
	fmt.Println(tabla)

	tabla2 := [5]int8{1, 2, 3, 4, 5}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}
	fmt.Println(matriz)

	//slices vectores dinamicos
	slices := []int{2, 5, 1}
	slices[2] = 5
	slices = append(slices, 9)
	fmt.Println(slices)
	variante2()
	variante3()
	variante4()
}
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d \n", len(elementos), cap(elementos))
}
func variante4() {
	elementos := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		elementos = append(elementos, i)
	}
	fmt.Println(elementos)
	fmt.Printf("largo %d, capacidad %d \n", len(elementos), cap(elementos))
}
