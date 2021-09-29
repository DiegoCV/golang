package main

import (
	"fmt"
	"strconv"
)

var estado bool
var hola string

func main() {
	if estado = true; estado {
		fmt.Println("Hola mundo." + hola)
	} else if hola = "nadar"; !estado {
		fmt.Println("Hola mundo. " + strconv.FormatBool(estado) + hola)
	}
	fmt.Println(hola)
	switch numero := 5; numero {
	case 1:
		fmt.Println("QUE? " + strconv.Itoa(numero))
	default:
		fmt.Println("QUE 2 ? " + strconv.Itoa(numero))
	}
}
