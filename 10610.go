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
	var thr int
	var zero bool
	var numbers []string
	var ans, N string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	num := strings.Split(N, "")

	for i := 0; i < len(num); i++ {
		if num[i] == "0" {
			zero = true
			numbers = append(numbers, num[i])
		} else {
			x, _ := strconv.Atoi(num[i])
			thr += x
			numbers = append(numbers, num[i])
		}
	}

	if (zero != true) || (thr%3 != 0) {
		fmt.Println(-1)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(numbers)))
		for i:=0;i<len(numbers);i++ {
			ans += numbers[i]
		}
		res, _ := strconv.Atoi(ans)
		fmt.Println(res)
	}
}
