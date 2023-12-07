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
)

const (
	height = 10
	width  = 10
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

func main() {
	a := NewUniverse()
	a.Seed()
	a.Show()
}
