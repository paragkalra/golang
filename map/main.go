package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	colors3["white"] = "#ffffff"
	fmt.Println(colors3)

	delete(colors3, "white")
	fmt.Println(colors3)

	value := colors["red"]
	fmt.Println(value)

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for", key, "is", value)
	}
}
