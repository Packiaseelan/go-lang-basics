package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	//multiple ways to create a map
	//var colors map[string]string
	// colors := make(map[string]string)
	// colors["White"] = "#ffffff"
	// colors["Black"] = "#000000"

	//delete(colors, "White")
	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
