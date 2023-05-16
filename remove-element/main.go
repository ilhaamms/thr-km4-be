package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	} else {
		return arr[:len(arr)-1]
	}
}

func main() {
	result := removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono", 100}, "right")
	fmt.Println(result)
}
