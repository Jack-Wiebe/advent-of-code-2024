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
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numSafe := 0

	for scanner.Scan() {
		line := scanner.Text()
		records := strings.Split(line, " ")

		nums := make([]int, len(records))
		for i, record := range records {
			val, _ := strconv.Atoi(record)
			nums[i] = val
		}

		if testSafe(nums) {
			numSafe++
			continue
		}

		for i := range nums {
			clone := make([]int, len(nums))
			copy(clone, nums)
			resliced := append(clone[:i], clone[i+1:]...)
			if testSafe(resliced) {
				numSafe++
				break
			}
		}
	}

	fmt.Println("Part 2:",numSafe)
}

func testSafe(nums []int) bool {
	dir := (nums[1] - nums[0]) > 0
	for i := 1; i < len(nums); i++ {
		dif := nums[i] - nums[i-1]
		//fmt.Println(nums)
		if(dif > 0 != dir){
			//fmt.Println("direction change")
			return false
		}
		if(dif == 0){
			//fmt.Println("no movement")
			return false
		}
		if(dif > 3 || -dif > 3){
			//fmt.Println("overload")
			return false
		}
	}

	return true
}


