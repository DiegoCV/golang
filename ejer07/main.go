package main

import "fmt"

var calculo func(int, int) int = func(i1, i2 int) int {
	return i1 + i2
}

func main() {
	fmt.Println("Hola mundo.")
	fmt.Printf("suma 5 + 7 %d \n", calculo(5, 7))
	calculo = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Printf("restar 5 - 7 %d \n", calculo(5, 7))
	fmt.Printf("a ber")
	gg := Tabla(1)
	fmt.Println(gg())
	fmt.Println(gg())
	fmt.Println(gg())
}

func Tabla(valor int) func() int {
	numero := valor
	var secuencia int
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
