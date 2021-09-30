package main

import (
	"fmt"
	"strings"
	"time"
)

func nombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}

func main() {
	go nombreLento("Diego Alirio")
	var x string
	fmt.Scanln(&x)
}
