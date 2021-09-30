package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	if h.esHombre {
		return "hombre"
	} else {
		return "mujer"
	}
}
func (h *hombre) estaVivo() bool {
	return h.vivo
}
func humanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un %s y estoy respirando\n", hu.sexo())
}

/**************/
type perro struct {
	respirando bool
	comeCarne  bool
	comiendo   bool
}

func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) esCarnivoro() bool {
	return p.comeCarne
}

func seresVivosViviendo(sv serVivo) {
	msg := "muerto"
	if sv.estaVivo() {
		msg = "vivo"
	}
	fmt.Printf("estoy %s\n", msg)
}

func main() {
	pepe := new(hombre)
	maria := new(mujer)
	pepe.esHombre = true
	pepe.vivo = true
	humanosRespirando(pepe)
	humanosRespirando(maria)
	seresVivosViviendo(pepe)
	seresVivosViviendo(maria)

}
