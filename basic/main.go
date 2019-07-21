package main

import (
	"fmt"
)

func main() {

	fmt.Println("Program begin")

	json := map[string]int{
		"nguyen": 6,
		"dinh":   4,
		"dung":   4,
	}

	for key, val := range json {
		fmt.Println(key, val)
	}

}
