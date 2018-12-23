package main

import "fmt"

func main() {

	// declaringMap()
	iterateOverMap()
}

// helper method - separate examples
func iterateOverMap() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"blue":  "#0000FF",
	}

	printMap(colors)

	fmt.Println("-------------------")

	colors["red"] = "red"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s --> \"%s\"\n", color, hex)
	}
}

// helper method - separate examples
func declaringMap() {
	// 1. declare map
	// zero value - empty map `map[]`
	// var colors map[string]string

	// 2. declare map - identical to #1
	// colors := make(map[string]string)

	// add item to map
	// colors["white"] = "#FFFFFF"

	// 3. declare map
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"blue":  "#0000FF",
	}

	fmt.Println(colors)

	// delete element from map
	delete(colors, "green")

	fmt.Println(colors)
}
