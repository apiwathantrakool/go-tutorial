package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["black"] = "#000000"
	delete(colors, "green")

	printMap(colors)
}

func printMap(c map[string]string) {
	for id, value := range c {
		fmt.Println("id is", id, ", and value is", value)
	}
}
