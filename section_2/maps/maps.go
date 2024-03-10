package main

import "fmt"

func main() {
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"

	delete(colors, "red")

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println("Hex code for", key, "is", val)
	}
}
