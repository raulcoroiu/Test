package main

import (
	"fmt"
)

func main() {
	//primul string reprezinta tipul chestii si cel de al 2-lea reprezinta tipul value-ului
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf654",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}

}
