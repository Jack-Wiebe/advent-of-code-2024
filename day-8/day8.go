package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file,_ := os.Open("input")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var input [][]string

	for scanner.Scan(){
		str := strings.Split(scanner.Text(), "")
		input = append(input, str)
		fmt.Println(str)
	}

	var node_types []string
	node_map := make(map[string][]Position)

	for row, line := range input{
		for col, cell := range line{
			if cell == "."{
				continue
			}
			if !contains(node_types, cell){
				node_types = append(node_types, cell)
			}else{
				node_map[cell] = append(node_map[cell], Position{x:row, y:col})
			}
		}
	}

	fmt.Println(node_types)
	fmt.Println(node_map)

}

type Position struct {
	x int
	y int
}

func contains[T comparable](input []T, value T) bool {
	for _, v := range input {
		if v == value {
			return true
		}
	}
	return false
}