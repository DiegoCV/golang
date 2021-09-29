package main

import "fmt"

func main() {
	paises := make(map[string]string)
	paises["colombia"] = "frutas y cafe"
	fmt.Println(paises)
	fmt.Println(paises["colombia"])
	paises["Argentina"] = "frio"

	gatos := make(map[string]string, 2)
	gatos["puma"] = "negro"
	fmt.Println(gatos["puma"])

	campeonato := map[string]int{
		"cucuta":    5,
		"santa_fe":  6,
		"argentina": 9}
	delete(campeonato, "santa_fe")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s tiene un puntaje: %d \n", equipo, puntaje)
	}
	puntaje, existe := campeonato["cucuta"]
	fmt.Printf("puntaje %d existe %t \n", puntaje, existe)
}
