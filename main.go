package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var arr []int
	filename := os.Args[1]
	info,_ := os.ReadFile(filename)
	split := strings.Split(string(info), "\n")
	for _, i := range split{
		num,_ := strconv.Atoi(i)
		arr = append(arr, num)
	}
	fmt.Println(average(arr))
	fmt.Println(median(arr))
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
	if len(a)%2 == 0 {
		half := (len(a)/2)-1
		median = a[half] + a[half+1] / 2
	} else {
		half := (len(a)/2)
		median = a[half]
	}

	return median
}