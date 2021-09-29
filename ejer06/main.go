package main

import "fmt"

func main() {
	fmt.Println("Hola mundo.")
	fmt.Println(uno(5))
	numero, flag := dos(4)
	fmt.Println(numero, flag)
	numero1, flag2, numero2 := tres(1)
	fmt.Println(numero1, flag2, numero2)
	fmt.Println("jajaja")
	calculo(1, 2, 3, 4, 5)
}

func uno(numero int) (z int) {
	z = numero * 2
	z -= 5
	z += 8
	return
}

func dos(numero int) (int, bool) {
	return numero * 2, (numero % 2) == 0
}

func tres(numero int) (int, bool, int) {
	return numero * 2, (numero % 2) == 0, 5
}

func calculo(numero ...int) {
	var total int
	for _, num := range numero {
		total += num
	}
	fmt.Println(total)
}
