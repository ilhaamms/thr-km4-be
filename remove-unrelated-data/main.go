package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	if _, ok := dataMap[key]; ok {
		delete(dataMap, key)
	}

	return dataMap
}

func main() {
	result := removeUnrelated(map[string]any{"run": "lari", "jump": "loncat", "swim": "berenang"}, "jump")
	fmt.Println(result)
}
