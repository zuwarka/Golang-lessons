// Измените следующую программу для обработки високосных годов. Код должен:
//
//    Генерировать случайный год вместо постоянного использования 2018;
//    Для февраля присвойте daysInMonth на 29 для високосных годов, и 28 для всех остальных. Можете использовать оператор if вместо блока case;
//    Используйте цикл for для генерации и отображения 10 случайных дат.

/*
package main

import (
    "fmt"
    "math/rand"
)

var era = "AD"

func main() {
    year := 2018
    month := rand.Intn(12) + 1
    daysInMonth := 31

    switch month {
    case 2:
        daysInMonth = 28
    case 4, 6, 9, 11:
        daysInMonth = 30
    }

    day := rand.Intn(daysInMonth) + 1
    fmt.Println(era, year, month, day)
}
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		year := rand.Intn(2100) + 1
		mounth := rand.Intn(12) + 1
		leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		daysInMonth := 31

		if leap && mounth == 2 {
			daysInMonth = 29
		} else if !leap && mounth == 2 {
			daysInMonth = 28
		}

		if !leap {
			switch mounth {
			case 2:
				daysInMonth = 28
			case 4, 6, 9, 11:
				daysInMonth = 30
			}
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, mounth, day)
	}
}
