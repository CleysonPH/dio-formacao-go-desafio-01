package main

import "fmt"

const pontoEbulicaoAguaEmKelvin = 373.2

func main() {
	pontoEbulicaoAguaEmCelsius := pontoEbulicaoAguaEmKelvin - 273
	fmt.Printf(
		"O ponto de ebulição da água em K é %g e em °C é %g",
		pontoEbulicaoAguaEmKelvin,
		pontoEbulicaoAguaEmCelsius,
	)
}
