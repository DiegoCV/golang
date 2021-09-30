package main

import (
	"fmt"
	"io/ioutil"
)

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./miArchivo.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	/*archivo, err := os.Open("./miArchivo.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		scanner = bufio.NewScanner(archivo)
		for scanner.ScanRawStrings(){
			texto := scanner.String
		}
	}*/
}

func main() {
	leoArchivo()
}
