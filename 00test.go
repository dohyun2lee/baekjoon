package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		scanner.Scan()
		stack := list.New()

		x := scanner.Text()
		X := strings.Split(x, ",")

		if len(X) == 1 {
			var tmp []string
			a := strings.Split(X[0], "")
			for i := 0; i < len(a); i++ {
				if a[i] != "[" && a[i] != "]" {
					tmp = append(tmp, a[i])
				}
			}
			if len(tmp) == 0 {
				stack = list.New()
			} else {
				stack.PushBack(strings.Join(tmp, ""))
			}
		} else if len(X) == 2 {
			first := strings.Split(X[0], "")
			last := strings.Split(X[len(X)-1], "")

			first = first[1:]
			last = last[:len(last)-1]

			stack.PushBack(strings.Join(first, ""))
			stack.PushBack(strings.Join(last, ""))
		} else {
			first := strings.Split(X[0], "")
			last := strings.Split(X[len(X)-1], "")

			first = first[1:]
			last = last[:len(last)-1]

			stack.PushBack(strings.Join(first, ""))
			for i := 1; i < len(X)-1; i++ {
				stack.PushBack(X[i])
			}
			stack.PushBack(strings.Join(last, ""))
		}

		for e := stack.Front(); e != nil; e = e.Next() {
			fmt.Fprint(writer, e.Value)
		}
		
	}
}
