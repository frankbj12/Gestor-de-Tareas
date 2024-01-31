package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64

	// Solicitar el primer número
	fmt.Print("Ingresa el primer número: ")
	fmt.Scanln(&num1)

	// Solicitar el segundo número
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scanln(&num2)

	// Mostrar opciones al usuario
	fmt.Println("Selecciona la operación:")
	fmt.Println("1. suma")
	fmt.Println("2. resta")
	fmt.Println("3. multiplicación")
	fmt.Println("4. división")

	// Leer la opción del usuario
	var opcion string
	fmt.Scanln(&opcion)

	// Realizar la operación seleccionada
	switch opcion {
	case "1":
		resultado := num1 + num2
		fmt.Println("Resultado:", resultado)
	case "2":
		resultado := num1 - num2
		fmt.Println("Resultado:", resultado)
	case "3":
		resultado := num1 * num2
		fmt.Println("Resultado:", resultado)
	case "4":
		if num2 != 0 {
			resultado := num1 / num2
			fmt.Println("Resultado:", resultado)
		} else {
			fmt.Println("Error: No se puede dividir por cero.")
		}
	default:
		fmt.Println("Opción no válida.")
	}
}
