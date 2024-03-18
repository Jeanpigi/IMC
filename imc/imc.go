// imc.go
package imc

import "math"

// CalcularIMC calcula el Índice de Masa Corporal basado en el peso (kg) y la estatura (m).
func CalcularIMC(peso, estatura float64) float64 {
	return peso / math.Pow(estatura, 2)
}

// CategoriaIMC determina la categoría de peso basada en el IMC.
func CategoriaIMC(imc float64) string {
	switch {
	case imc < 18.5:
		return "Bajo peso"
	case imc < 25.0:
		return "Normal"
	case imc < 30.0:
		return "Sobre peso"
	case imc < 35.0:
		return "Obesidad"
	case imc < 40.0:
		return "Obesidad grave"
	default:
		return "Obesidad mórbida"
	}
}
