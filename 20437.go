package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T, K int
	var S string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &S)
		fmt.Fscanln(reader, &K)
		var test map[string]int = make(map[string]int)
		var cnt, max, min, tmp, start, max_cnt, max_tmp int
		var ans_min, ans_max string
		var set []string
		chk := false
		start = -1

		s := strings.Split(S, "")

		for j := 0; j < len(s); j++ {
			test[s[i]]++
		}

		for key, val := range test {
			if val >= K {
				chk = true
				set = append(set, key)
			}
		}

		if chk {
			for {
				for j := 0; j < len(s); j++ {
					if start < 0 && s[j] == set[tmp] {
						start = j
					}
					if s[j] == set[tmp] {
						if max_cnt != K-1 {
							max_cnt++
						} else {
							if max_tmp <= j {
								max_tmp = j
							}
						}
						cnt++
					}
					if cnt == K && min == 0 {
						min = j
					}
				}
				if tmp == len(s)-1 {
					break
				}

			}
		}
	}
}
