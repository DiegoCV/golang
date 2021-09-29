package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2 int
var leyenda string

func main() {
	fmt.Println("Hola mundo.")
	fmt.Scanln(&numero1)
	fmt.Println("Hola mundo.")
	fmt.Scanln(&numero2)
	fmt.Println("dsfsd")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	fmt.Printf("%d", numero1+numero2)
	fmt.Println(leyenda)
}
