// Программа должна построить две таблицы. В первой таблице два столбца, в первом значится температура по Цельсию °C,
// а во втором — по Фаренгейту °F. Значения должны быть от 40° C до 100° C шагами в 5°.
// Для заполнения столбцов требуется использовать методы конвертации, описанные в уроке о методах.
//
// После заполнения одной таблицы заполните вторую таким образом, чтобы столбцы были инвертированы.
// То есть конвертация должна проводиться из градусов по Фаренгейту в градусы по Цельсию.
//
// Код, что вы напишите для создания таблицы, в будущем можно будет использовать вновь, уже для других программ,
// содержимое которых нужно отобразить в таблице с двумя столбцами.
// Используйте функции для разделения кода который создает таблицы от кода для вычисления значений температуры каждой строки.
//
// Реализуйте функцию drawTable, что принимает функцию первого класса в качестве параметра и вызывает ее
// для получения данных каждой созданной строки.
// Результатом передачи другой функции к drawTable должны быть другие отображаемые данные.
package main

import "fmt"

type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

// по сути это такая виртуальнгая функция как в си, которая потом переопределяется
type getRowFn func(row int) (string, string)

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)

	for i := 0; i < rows; i++ {
		cell1, cell2 := getRow(i)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := celsius(row)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", 15, ctof)
	fmt.Println()
	drawTable("°F", "°C", 15, ftoc)
}
