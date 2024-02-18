package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N, M, cnt, Distance int
	var s, ans string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var S [][]string

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &s)
		S = append(S, strings.Split(s, ""))
	}

	for i := 0; i < M; i++ {
		var alphabet map[string]int = make(map[string]int)
		for j := 0; j < N; j++ {
			alphabet[S[j][i]]++
		}
		ans += pick(alphabet)
	}

	tmp := strings.Split(ans, "")

	for i := 0; i < N; i++ {
		cnt = 0
		for j := 0; j < M; j++ {
			if tmp[j] != S[i][j] {
				cnt++
			}
		}
		Distance += cnt
	}

	fmt.Println(ans)
	fmt.Println(Distance)
}

func pick(alphabet map[string]int) string {
	var max int
	var spell string
	for key, val := range alphabet {
		if val >= max {
			if max == val {
				if spell >= key {
					spell = key
				}
			} else {
				max = val
				spell = key
			}
		}
	}

	return spell
}
