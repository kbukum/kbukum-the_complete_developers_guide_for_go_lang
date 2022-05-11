package main

import "fmt"

/**
Map <-> Hash <-> Object <-> Dict
- key is in the same type
- value is in the same type
 **/
func main() {

	fmt.Println("# Creating string, string map")
	colorsHexMap := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println(colorsHexMap)

	fmt.Println()
	fmt.Println("# Creating empty map")
	var colorsHexMap2 map[string]string
	fmt.Println(colorsHexMap2)

	fmt.Println("# Using for loop on the map to get each entry!")
	for key, value := range colorsHexMap {
		fmt.Printf("%s=%s\n", key, value)
	}

	fmt.Println("# Adding element to the map!")
	colorsHexMap["white"] = "#FFFFFF"
	for key, value := range colorsHexMap {
		fmt.Printf("%s=%s\n", key, value)
	}

	fmt.Println("# Removing element to the map!")
	delete(colorsHexMap, "red")
	for key, value := range colorsHexMap {
		fmt.Printf("%s=%s\n", key, value)
	}
}
