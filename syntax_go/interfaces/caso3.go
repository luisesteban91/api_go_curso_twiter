package main

import "fmt"

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *perro) respirar() { p.respirando = true }
func (p *perro) comer()    { p.comiendo = true }
func (p *perro) EsCarnivoro() bool {
	return p.carnivoro
}

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Printf("Soy un animal, y ya estoy respirando \n", an.EsCarnivoro())
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

func main() {
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d", totalCarnivoros)
}
