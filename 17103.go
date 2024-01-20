package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

var primenumber []int
var num []int

func main() {
	var T, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var numCopy []int

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &N)
		num = append(num, N)
		numCopy = append(numCopy, N)
	}

	sort.Ints(numCopy)

	for i:=2;i<numCopy[T-1];i++ {
		if primeNum(i) == true {
			primenumber = append(primenumber, i)
		}
	}

	for i:=0;i<T;i++ {
		cnt := 0
		x := spcPrimeNum(num[i])	
		for j:=0;j<=x;j++ {
			for k:=j;k<=x;k++ {
				if primenumber[j] + primenumber[k] == num[i] {
					cnt++
				}
			}
		}
		fmt.Fprintln(writer, cnt)
	}
}

func primeNum(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func spcPrimeNum(x int) int {
	var a int
	for i:=0;i<len(primenumber);i++ {
		if primenumber[i] > x {
			break
		}
		a = i
	}
	return a
}