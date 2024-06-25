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
	if len(os.Args) != 2 {
		fmt.Println("go run main.go data.txt")
		return
	}
	filename := os.Args[1]
	info, _ := os.ReadFile(filename)
	split := strings.Split(string(info), "\n")
	for _, i := range split {
		num, _ := strconv.Atoi(i)
		arr = append(arr, num)
	}
	// if there is a newline
	if arr[len(arr)-1] == 0 { 
	arr = arr[:len(arr)-1]
	}

	fmt.Println("Average:",average(arr))
	fmt.Println("Median:",median(arr))
	fmt.Println("Variance:",variance(arr))
	fmt.Println("Standard Deviation:",deviation(arr))
}

func average(a []int) int {// 1 2
	len := len(a)
	var n int
	for _, i := range a {
		n += i
	}
	new := float64(n)
	new /= float64(len)
	me := int(math.Round(new))
	return me
}

func median(a []int) int {
	b := sort(a)
	var median int
	med := make([]int, 2)
	if len(a)%2 == 0 { // 1 2   3 4   5 6
		half := (len(a) / 2) - 1 // 1 2 3 4
		med[0], med[1] = b[half], b[half+1]
		median = average(med)
	} else {
		half := (len(a) / 2)// 1 2 3 4 5
		median = b[half]
	}
	return median
}

func variance(a []int) int {
	x := average2(a)
	var res float64
	var variance int
	for _, i := range a {
		res += power((float64(i) - x), 2)
	}
	variance = int(math.Round(res / float64(len(a))))
	return variance
}

func power(n, y float64) float64 {
	if n < 0 {
		n = -n
	}
	if y == 0 {
		return 1
	}
	for i := 1; float64(i) < y; i++ {
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

func sort(a []int) []int{
	for range a {
	for i := range a[:len(a)-1] {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
}
return a
}

//average with decimals
func average2(a []int) float64 {// 1 2
	len := len(a)
	var n int
	for _, i := range a {
		n += i
	}
	new := float64(n)
	new /= float64(len)
	return new
}