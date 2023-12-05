//
// Допишите Листинг 8 для отображения всех шахматных фигур на их стартовых позициях,
//используя символы kqrbnp для черных фигур в верхней части доски,
//а также символы в верхнем регистре KQRBNP для белых фигур в нижней части доски;
//Напишите функцию для отображения доски;
//Вместо строк, используйте [8][8]rune для доски.
//Помните, что литералы rune должны быть окружены одинарными кавычками
//и могут выводиться на экран через специальный символ %c.

package main

import "fmt"

func display(board [8][8]rune) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}

func main() {
	var board [8][8]rune

	// черные фигуры
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	// белые фигуры
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	// пешки
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}

	display(board)
}
