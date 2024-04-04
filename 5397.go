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
	var S string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		l := list.New()

		fmt.Fscanln(reader, &S)
		s := strings.Split(S, "")

		l.PushBack("*")
		var cursor = l.Back()

		for j := 0; j < len(s); j++ {
			switch s[j] {
			case "<":
				if cursor.Prev() != nil {
					l.MoveBefore(cursor, cursor.Prev())
				}
			case ">":
				if cursor.Next() != nil {
					l.MoveAfter(cursor, cursor.Next())
				}
			case "-":
				if cursor.Prev() != nil {
					l.Remove(cursor.Prev())
				}
			default:
				l.InsertBefore(s[j], cursor)
			}
		}

		l.Remove(cursor)

		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Fprint(writer, e.Value)
		}
		fmt.Fprintln(writer)
	}
}
