package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var A, B, x int
	arr := list.New()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A, &B)

	sum := A + B
	Sum := strings.Split(strconv.Itoa(sum), "")
	
	for i := len(Sum)-1; i >= 0; i-- {
		num, _ := strconv.Atoi(Sum[i])
		switch num + x {
		case 0:
			arr.PushBack(0)
			x = 0
		case 1:
			arr.PushBack(1)
			x = 0
		case 2:
			arr.PushBack(0)
			x = 1
		case 3:
			arr.PushBack(1)
			x = 1
		}
	}

	if x == 1 {
		arr.PushBack(1)
	}

	for e := arr.Back(); e != nil; e = e.Prev() {
		fmt.Fprint(writer, e.Value)
	}
}
