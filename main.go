package main

import (
	"fmt"

	"github.com/rcotilla/gocalc/pkg"
)

func main() {
	fmt.Println("Calculadora basica de operaciones matematicas.")
	fmt.Println("Ingrese la operacion a realizar (ej: `1+2*(3-1)` ): ")

	var op string
	_, err := fmt.Scan(&op)
	if err != nil {
		fmt.Println("Error al leer la operacion.")
		return
	}

	fmt.Println(pkg.Calculate(op))
	fmt.Println("Presione Enter para salir.")
	fmt.Scanln()
}
