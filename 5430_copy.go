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
		var flip int = 1
		var N, err int
		var cmd string

		fmt.Fscanln(reader, &cmd)
		fmt.Fscanln(reader, &N)
		scanner.Scan()

		stack := list.New()
		x := scanner.Text()
		X := strings.Split(x, ",")
		command := strings.Split(cmd, "")

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

		for j := 0; j < len(command); j++ {
			if command[j] == "R" {
				flip *= -1
			} else {
				if stack.Len() == 0 {
					err = 1
					break
				} else {
					if flip > 0 {
						stack.Remove(stack.Front())
					} else {
						stack.Remove(stack.Back())
					}
				}
			}
		}

		if err > 0 {
			fmt.Fprintln(writer, "error")
		} else {
			if flip > 0 {
				fmt.Fprintf(writer, "[")
				for e := stack.Front(); e != nil; e = e.Next() {
					if e.Next() == nil {
						fmt.Fprint(writer, e.Value)
					} else {
						fmt.Fprint(writer, e.Value)
						fmt.Fprint(writer, ",")
					}
				}
				fmt.Fprintln(writer, "]")
			} else {
				fmt.Fprintf(writer, "[")
				for e := stack.Back(); e != nil; e = e.Prev() {
					if e.Prev() == nil {
						fmt.Fprint(writer, e.Value)
					} else {
						fmt.Fprint(writer, e.Value)
						fmt.Fprint(writer, ",")
					}
				}
				fmt.Fprintln(writer, "]")
			}
		}
	}
}
