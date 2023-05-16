package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {

	result := howManyElements([]any{"1", "2", "Joko", false, 5.2, 6, 7, 8, 9})
	fmt.Println(result)

}
