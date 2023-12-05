// Используйте calibrate с функцией fakeSensor из Листинга 2 для создания новой функции sensor. Вызовите новую функцию
// sensor несколько раз, обратите внимание, что оригинальная функция fakeSensor по-прежнему вызывается каждый раз,
// выдавая в результате случайные значения.
package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	result := kelvin(rand.Intn(151) + 150)
	fmt.Printf("я фейк, %f", result)
	return result
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(fakeSensor, offset)

	for count := 0; count < 10; count++ {
		fmt.Println(sensor())
	}
}
