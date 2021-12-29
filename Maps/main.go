package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
	}
	fmt.Println(colors) //output-> map[green:#00FF00 red:#ff0000]

	//
	var colors2 map[string]string
	fmt.Println(colors2) //output-> map[]

	//
	colors3 := make(map[int]string)
	fmt.Println(colors3) //output-> map[]

	colors3[10] = "white"
	fmt.Println(colors3) //output-> map[10:white]

	//delete
	delete(colors3, 10)
	fmt.Println(colors3) //output-> map[]

	//iterating over map
	printMap(colors) //output-> Hex code for green is #00FF00...
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
