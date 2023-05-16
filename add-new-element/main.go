package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	} else {
		return data
	}

}

func main() {
	result := AddElement([]int{1, 2, 3, 4, 5}, 6, "down")
	fmt.Println(result)
}
