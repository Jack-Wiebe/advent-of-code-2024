package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	dif := 0
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

	sort.Ints(right_list)
	sort.Ints(left_list)

	for i := 0; i < len(left_list); i++ {
		dif += absInt(left_list[i] - right_list[i])
	}

	fmt.Println(dif)

	part2()

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}