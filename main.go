package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jeanpigi/IMC/imc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Calculadora de Índice de Masa Corporal (IMC)")
	fmt.Print("Introduce tu peso en kilogramos: ")
	scanner.Scan()
	peso, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Error al leer el peso:", err)
		return
	}

	fmt.Print("Introduce tu estatura en metros: ")
	scanner.Scan()
	estatura, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Error al leer la estatura:", err)
		return
	}

	imcResultado := imc.CalcularIMC(peso, estatura)
	categoria := imc.CategoriaIMC(imcResultado)

	fmt.Printf("Tu Índice de Masa Corporal (IMC) es: %.2f. Categoría: %s\n", imcResultado, categoria)
}
