// Напишите функцию celsiusToKelvin, что использует типы celsius и kelvin, определенные в Листинге 4.
// Используйте ее для конвертации 127° C, температуры освещенной солнцем поверхности луны, в градусы Кельвина.
package main

import "fmt"

type celsius float64
type kelvin float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 36.7
	k := celsiusToKelvin(c)
	fmt.Println(k)
}
