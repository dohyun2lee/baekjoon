package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S, x, cmd string
	var M int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &S)
	fmt.Fscanln(reader, &M)

	l := list.New()
	s := strings.Split(S, "")

	for i := 0; i < len(s); i++ {
		l.PushBack(s[i])
	}

	l.PushBack("|")
	var cursor = l.Back()

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case "L":
			if cursor.Prev() != nil {
				l.MoveBefore(cursor, cursor.Prev())
			}
		case "D":
			if cursor.Next() != nil {
				l.MoveAfter(cursor, cursor.Next())
			}
		case "B":
			if cursor.Prev() != nil {
				l.Remove(cursor.Prev())
			}
		case "P":
			l.InsertBefore(x, cursor)
		}
	}

	l.Remove(cursor)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Fprint(writer, e.Value)
	}
}
