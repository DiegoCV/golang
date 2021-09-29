package main

import "fmt"

func main() {

	i := 1
	for i < 10 {
		fmt.Println("Hola mundo. ", i)
		i++
	}
	for j := i; j < 20; j++ {
		fmt.Println("Hola mundo. ", j)
	}
	fmt.Println("infinito ")
	numero := 0
	for {
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
		fmt.Println("EEEE ", numero)
	}
	k := 5
MARCA:
	k++
	fmt.Printf("\n aqui k %d", k)
	if k == 6 {
		goto MARCA
	}
}
