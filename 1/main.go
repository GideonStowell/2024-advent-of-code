package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
	"strings"
	"strconv"
	"sort"

)
func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var right []int
	var left []int
	for scanner.Scan() {
		split := strings.Split(scanner.Text(),"   ")
		r_val, _ := strconv.Atoi(split[1])
		l_val, _ := strconv.Atoi(split[0])
		right = append(right, r_val)
		left = append(left, l_val)
	}

	sort.Ints(left)
	sort.Ints(right)
	printSlices(left)
	printSlices(right)

	if len(right) != len(left) {
		panic("lists are not equal")
	}


	var sum int

	for index, _ := range right {
		// fmt.Println(math.Abs(left[index] - right[index]))
		sum = sum + difference(left[index], right[index])
	}
	fmt.Println("sum")
	fmt.Println(sum)

	// # part2

	right_freq := make(map[int]int)
	// var freq int
	for _, item := range right {


		right_freq[item] = right_freq[item]+1
	}
	printMap(right_freq)
	score := 0
	for _, item := range left {
		multiplier := right_freq[item]
		if multiplier > 0 {
			score = score + (item * multiplier)
		}
	}
	fmt.Println(score)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printMap(m map[int]int){
	for key, value := range m {
    fmt.Printf("%v => %v\n", key, value)
}
}

func difference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}