package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var N, M int
	var name string
	var pokemon map[int]string = make(map[int]string, N)
	var pokemon2 map[string]int = make(map[string]int, N)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i:=1;i<=N;i++ {
		fmt.Fscanln(reader, &name)
		pokemon[i] = name
		pokemon2[name] = i
	}

	for j:=1;j<=M;j++ {
		var a string
		fmt.Fscanln(reader, &a)
		A, err := strconv.Atoi(a)

		if err == nil {
			fmt.Fprintln(writer, pokemon[A])
		} else {
			fmt.Fprintln(writer, pokemon2[a])
		}
	}
}
