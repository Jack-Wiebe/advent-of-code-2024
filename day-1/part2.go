package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left_list []int
	var right_list []int

	for scanner.Scan(){
		//fmt.Println(scanner.Text())
		str := scanner.Text()
		result := strings.Fields(str)

		int0, _ := strconv.Atoi(result[0])
		int1, _ := strconv.Atoi(result[1])

		left_list = append(left_list, int0)
		right_list = append(right_list, int1)

	}

	foo := 0

	for i := 0; i < len(left_list); i++ {
		mult := count(right_list, left_list[i])
		foo += mult * left_list[i]
	}

	fmt.Println(foo)

}

func count(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

