package main

import "fmt"

func exponente(base int, exp int) int {
	if exp == 0 {
		return 1
	} else {
		return base * exponente(base, exp-1)
	}
}

func main() {
	fmt.Println("Hola mundo.", exponente(2, 20))
}
