package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input")

	if err != nil{
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var report [][]string

	for scanner.Scan(){
		str := scanner.Text()
		line := strings.Split(str, " ")
		//fmt.Println(line)
		report = append(report, line)
	}
	fmt.Println()

	count := 0

	for i := 0; i < len(report); i++{
		dir := 0
		safe := true

		int0, _ := strconv.Atoi(report[i][0])
		int1, _ := strconv.Atoi(report[i][1])

		if  int0 == int1{
			dir = 0
			safe = false
		}else if int0 > int1{
			dir = -1
		}else{
			dir = 1
		}

		//fmt.Println(report[i])
		//fmt.Println("calibrate direction:", report[i][0], report[i][1], "direction:", dir)

		for j := 1; j < len(report[i]); j++{

			int0, _ := strconv.Atoi( report[i][j])
			int1, _ := strconv.Atoi(report[i][j - 1])

			dif := int0 - int1
			//fmt.Println(dif)

			dif = dif * dir

			if(dif < 0){
				//fmt.Println("direction change")
				safe = false
			}
			if(dif == 0){
				//fmt.Println("no movement")
				safe = false
			}
			if(dif > 3){
				//fmt.Println("overload")
				safe = false
			}
		}
		//fmt.Println("safe: ", safe)
		if safe{
			count++
		}

	}

	fmt.Println("part 1:", count)

	part2()
}