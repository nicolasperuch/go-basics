package main

import "fmt"

type rgb struct {
	red int
	green int
	blue int
}

func main() {

	hexadecimalColors := map[string] string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	//another way to create a map
	rgbColors := make(map[string] rgb)

	rgbColors["red"] = rgb {255, 0, 0}
	rgbColors["green"] = rgb {0, 255, 0}
	rgbColors["blue"] = rgb {0, 0, 255}
	rgbColors["black"] = rgb {0, 0, 0}

	delete(rgbColors, "black")

	fmt.Println("Hexademical colors:", hexadecimalColors)
	fmt.Println("Rgb colors:", rgbColors)

	printMap(hexadecimalColors)
}

func printMap(c map[string] string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}