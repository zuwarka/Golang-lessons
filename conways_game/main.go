// Правила Игры в «Жизнь» следующие:
//
//    Живая клетка, у которой менее двух живых соседей, умирает;
//    Живая клетка, у которой два или три живых соседа, живет до следующего поколения;
//    Живая клетка, у которой более трех живых соседей, умирает;
//    Мертвая клетка, у которой есть ровно три живых соседа, оживает.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	height    = 10
	width     = 10
	countShow = 300
)

type Universe [][]bool

func NewUniverse() Universe {
	newUniverse := make(Universe, width)
	for k, _ := range newUniverse {
		newUniverse[k] = make([]bool, width)
	}
	return newUniverse
}

func (u Universe) Show() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u[y][x] == true {
				fmt.Print("* ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Print("\n")
	}
}

func (u Universe) Seed() {
	for y := 0; y < rand.Intn(height); y++ {
		for x := 0; x < rand.Intn(width); x++ {
			u[y][x] = true
		}
		fmt.Print("\n")
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.Next(x, y)
		}
	}
}

// CheckAliveUniverse функция, которая проверяет состояние Вселенной на все false
func (u Universe) CheckAliveUniverse() bool {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u[y][x] == true {
				return true
			}
		}
	}

	return false
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < countShow; i++ {
		Step(a, b)
		if !a.CheckAliveUniverse() { // если все false - останавливаем цикл
			break
		}
		a.Show()
		fmt.Print("\x0c")
		time.Sleep(time.Second)
		a, b = b, a
	}
}
