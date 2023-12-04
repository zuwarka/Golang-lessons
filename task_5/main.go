// Напишите программу с типами celsius, fahrenheit и kelvin и методами для конвертации из одного типа температуры в другой.

package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// методы для celsius
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// методы для kelvin
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Print(k, "° K is ", c, "° C")
}

// методы для fahrenheit
