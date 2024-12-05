package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func test() {

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
		fmt.Println(line)
		report = append(report, line)
	}
	fmt.Println()

	count := 0
	for i := 0; i < len(report); i++{

		index := reportSafe(report[i])

		if(index == -1){
			fmt.Println("safe")
			count++
		}else{
			fmt.Println(report[i])
			test := splice(report[i],index)
			fmt.Println(test)

			index = reportSafe(test)
			if(index == -1){
				fmt.Println("safe")
				count++
			}
			fmt.Println(count)
			fmt.Println()
		}
	}

	fmt.Println(count)
}

func reportSafe(report []string) int {
	dir := 0

	val0 := report[0]
	val1 := report[1]
	int0, err0 := strconv.Atoi(val0)
	int1, err1 := strconv.Atoi(val1)

	if err0 != nil || err1 != nil{
		fmt.Println("FAIL")
	}

	if  int0 == int1{
		fmt.Println("start fail")
		return 0
	}else if int0 > int1{
		dir = -1
	}else{
		dir = 1
	}


	for j := 1; j < len(report); j++{

		val0 := report[j]
		val1 := report[j - 1]
		int0, err0 := strconv.Atoi(val0)
		int1, err1 := strconv.Atoi(val1)

		if err0 != nil || err1 != nil{
			fmt.Println("FAIL")
			return -1
		}

		dif := int0 - int1

		dif = dif * dir

		if(dif == 0){
			fmt.Println("no movement")
			if(j == 2){
				return 0
			}else{
				return j
			}
		}else if(dif < 0){
			fmt.Println("direction change")

			if(j+1 < len(report) && report[j + 1] == report[j - 1]){
				fmt.Println("sandwich")
				if(dir > 0){
					fmt.Println("rising")
					if(report[j] < report[j + 1]){
						return j + 1
					}else{
						return j - 1
					}
				}else{
					fmt.Println("falling")
					if(report[j] < report[j - 1]){
						return j + 1
					}else{
						return j - 1
					}
				}



			}else if(j == 2){
				return j-2
			}else{
				return j
			}
		}else if(dif > 3){
			fmt.Println("overload")
			if(j == 1){
				return 0
			}else{
				return j
			}
		}
	}
	return -1
}

//5,8,4,3,1

func splice(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	test := append(slice[:index], slice[index+1:]...)
	return test
}