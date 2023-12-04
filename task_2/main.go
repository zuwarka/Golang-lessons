// Напишите программу для вывода каждого байта (символ ASCII) слова «shalom», по одному символу на строку.

package main

import "fmt"

func main() {
	shalom := "shalom"

	for i := 0; i < len(shalom); i++ {
		c := shalom[i]
		fmt.Printf("%c\n", c)
	}
}
