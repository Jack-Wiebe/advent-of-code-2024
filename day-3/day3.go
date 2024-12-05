package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	pattern  := `mul[(](\d+)[,](\d+)[)]`
	sum := 0

	for scanner.Scan(){
		input += scanner.Text()
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches{

		num0,_ := strconv.Atoi(match[1])
		num1,_ := strconv.Atoi(match[2])
		sum += (num0 * num1)
	}

	fmt.Println("sum of matches in part 1 is:", sum)


	part2()
}