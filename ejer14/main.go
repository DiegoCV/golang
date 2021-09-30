package main

import (
	"log"
)

func main() {
	/*archivo := "prueba.txt"
	f, err := os.Open(archivo)
	defer f.Close()
	if err != nil {
		fmt.Println("Hola mundo. Error")
		os.Exit(1)
	}*/
	ejemploPanic()
}
func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("AQUI NADA \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de uno")
	}
}
