// Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты.
//Используйте программу для изменения названий планет Марс, Уран и Нептун.
//
//В первой итерации может использоваться функция terraform,
//но в конечной реализации должен быть введен тип Planets с методом terraform, похожим на sort.StringSlice.

package main

import "fmt"

type Planets []string

// метод типа Planets
func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "Новый " + planets[i]
	}
}

func main() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
}
