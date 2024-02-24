package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N, K, cnt int
	var S string
	var hamburger map[int]bool = make(map[int]bool)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)

	fmt.Fscanln(reader, &S)

	hp := strings.Split(S, "")

	for i:=0;i<N;i++ {
		if hp[i] == "H" {
			hamburger[i] = true
		}
	}

	for i:=0;i<N;i++ {
		if hp[i] == "P" {
			if i-K < 0 {
				for j:=0;j<=i+K;j++ {
					if hamburger[j] == true {
						cnt++
						hamburger[j] = false
						break
					}
				}
			} else {
				for j:=i-K;j<=i+K;j++ {
					if hamburger[j] == true {
						cnt++
						hamburger[j] = false
						break
					}
				}
			}
		}
	}

	fmt.Println(cnt)
}
