package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T, place int
	var S string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		place = 0
		var ans []string
		var tmp []string
		fmt.Fscanln(reader, &S)

		s := strings.Split(S, "")

		for j := 0; j < len(s); j++ {
			if s[j] == "<" {
				if place != 0 {
					place--
				}
			} else if s[j] == ">" {
				if place != j {
					place++
				}
			} else if s[j] == "-" {
				tmp = ans[:place]
				
			} else {
				
			}
			p++
		}
	}
}
