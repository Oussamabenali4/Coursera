package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a name: ")
	name := make([]byte, 20)
	n, err := os.Stdin.Read(name)
	if err != nil {
		fmt.Println("Error while  reading the name")
		return
	}
	fmt.Println("Enter a date")
	add := make([]byte, 1024)
	m, err := os.Stdin.Read(add)
	if err != nil {
		fmt.Println("Err while reading the Address ")
	}
	name = name[:n-1]
	add = add[:m-1]

	var Map = make(map[string]string)
	Map[string(name)] = string(add)

	j, err := json.Marshal(Map)
	if err != nil {
		fmt.Println("Error formating the map into json")
	}
	fmt.Println(string(j))
}
