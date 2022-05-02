package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}

	colors["white"] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colors["white"])

	delete(colors, "white")
	fmt.Println(colors)

	print(colors)
}

func print(colors map[string]string) {
	for key, value := range colors {
		fmt.Println(key, value)
	}
}
