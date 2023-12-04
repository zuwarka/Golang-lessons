// Расшифруйте цитату Юлия Цезаря: L fdph, L vdz, L frqtxhuhg.
//
//Ваша программа должна будет сдвинуть буквы верхнего и нижнего регистра на -3.
//Помните, что ‘a’ становится ‘x’, ‘b’ становится ‘y’, а ‘c’ становится ‘z’.
//То же самое происходит с буквами верхнего регистра.

package main

import "fmt"

func main() {
	msg := "L fdph, L vdz, L frqtxhuhg"

	for _, k := range msg {
		if k >= 'a' && k <= 'z' {
			k -= 3
			if k < 'a' {
				k += 26
			}
		} else if k >= 'A' && k <= 'Z' {
			k -= 3
			if k < 'A' {
				k += 26
			}
		}

		fmt.Printf("%c", k)
	}
}
