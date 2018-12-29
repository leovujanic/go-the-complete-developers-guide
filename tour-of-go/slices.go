package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"strings"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for y, row := range p {
		for x := range row {
			row[x] = uint8(x * y)
		}
	}

	return p
}

func grid(dx, dy int) [][]byte {
	g := make([][]byte, dy)

	for i := range g {
		g[i] = make([]byte, dx)
	}

	for _, row := range g {
		for x := range row {
			row[x] = '*'
		}
	}

	return g
}

func main() {
	pic.Show(Pic)

	g := grid(20, 20)
	for _, row := range g {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Printf("\n")
	}

	fmt.Println(strings.Repeat("-", len(g)*2))

	for i := 0; i < len(g); i += 1 {
		for j := 0; j < len(g[i]); j += 1 {
			fmt.Printf("%c ", g[i][j])
		}
		fmt.Printf("\n")
	}

}
