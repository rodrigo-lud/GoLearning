package main

import "fmt"

func main() {
	colorsMap1 := map[string]string{
		"Black":  "#000000",
		"Blue":   "#0000FF",
		"Gray":   "#808080",
		"Green":  "#008000",
		"Purple": "800080",
		"Red":    "FF0000",
		"White":  "#FFFFFF",
	}
	fmt.Println("1st Map of colors:\n", colorsMap1)

	colorsMap2 := make(map[string]string) // or colorsMap2 := map[string]string{}
	colorsMap2["Black"] = "#000000"
	colorsMap2["White"] = "#FFFFFF"
	fmt.Println("2st Map of colors:\n", colorsMap2)

	delete(colorsMap2, "Black")
	fmt.Println("2st Map of colors after delete:\n", colorsMap2)

	printMap(colorsMap1)
}

func printMap(m map[string]string) {
	for i1, i2 := range m {
		fmt.Println("The hex code for", i1, "is: ", i2)
	}
}
