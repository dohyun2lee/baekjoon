package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, q int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var questionWords []string = make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &questionWords[i])
		if questionWords[i] == "?" {
			q = i
		}
	}

	if q == 0 {
		end := questionWords[1][0]
		fmt.Fscanln(reader, &M)
		var answerWords []string = make([]string, M)

		for i := 0; i < M; i++ {
			fmt.Fscanln(reader, &answerWords[i])
			if answerWords[i][len(answerWords[i])-1] == end {
				fmt.Fprintln(writer, answerWords[i])
			}
		}
	} else if q == N-1 {
		start := questionWords[N-2][len(questionWords[N-2])-1]
		fmt.Fscanln(reader, &M)
		var answerWords []string = make([]string, M)

		for i := 0; i < M; i++ {
			fmt.Fscanln(reader, &answerWords[i])
			if answerWords[i][0] == start {
				fmt.Fprintln(writer, answerWords[i])
			}
		}
	} else {
		start := questionWords[q-1][len(questionWords[q-1])-1]
		end := questionWords[q+1][0]
		fmt.Fscanln(reader, &M)
		var answerWords []string = make([]string, M)

		for i := 0; i < M; i++ {
			fmt.Fscanln(reader, &answerWords[i])
			if answerWords[i][len(answerWords[i])-1] == end && answerWords[i][0] == start {
				fmt.Fprintln(writer, answerWords[i])
			}
		}
	}
}
