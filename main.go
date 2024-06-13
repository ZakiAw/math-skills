package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arr []int
	filename := os.Args[1]
	info, _ := os.ReadFile(filename)
	split := strings.Split(string(info), "\n")
	for _, i := range split {
		num, _ := strconv.Atoi(i)
		arr = append(arr, num)
	}
	fmt.Println(average(arr))
	fmt.Println(median(arr))
	fmt.Println(variance(arr))
	fmt.Println(deviation(arr))
}

func average(a []int) int {
	len := len(a)
	var n int
	for _, i := range a {
		n += i
	}
	if n%len != 0 {
		n /= len
		n += 1
		return n
	}
	n /= len
	return n
}

func median(a []int) int {
	var median int
	med := make([]int, 2)
	if len(a)%2 == 0 { // 1 2   3 4   5 6
		half := (len(a) / 2) - 1 //
		med[0], med[1] = a[half], a[half+1]
		median = average(med)
	} else {
		half := (len(a) / 2)
		median = a[half]
	}
	return median
}

func variance(a []int) int {
	x := average(a)
	var variance, res int
	for _, i := range a {
		res += power((i - x), 2)
	}
	variance = res / len(a)
	return variance
}

func power(n, y int) int {

	if n < 0 {
		n = -n
	}
	if y == 0 {
		return 1
	}
	for i := 1; i < y; i++ {
		n *= n
	}

	return n
}

func deviation(a []int) int {
	i := variance(a)
	b := math.Sqrt(float64(i))
	c := math.Round(b)
	return int(c)
}
