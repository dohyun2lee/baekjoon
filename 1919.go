package main 

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var A, B string
	var cnt int
	var ALetter map[string]int = make(map[string]int)
	var ATrue map[string]bool = make(map[string]bool)
	var BLetter map[string]int = make(map[string]int)
	var BTrue map[string]bool = make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)

	a := strings.Split(A, "")
	b := strings.Split(B, "")

	for i:=0;i<len(a);i++ {
		ALetter[a[i]]++
		ATrue[a[i]] = true
	}
	for i:=0;i<len(b);i++ {
		BLetter[b[i]]++
		BTrue[b[i]] = true
	}

	for key,val := range ALetter {
		if ATrue[key] && BTrue[key] {
			if val - BLetter[key] >= 0 {
				cnt += val - BLetter[key]
			} else {
				cnt += BLetter[key] - val
			}
			ATrue[key] = false
			BTrue[key] = false
		} else if ATrue[key] {
			cnt += val
			ATrue[key] = false
		}
	}

	for key,val := range BLetter {
		if ATrue[key] && BTrue[key] {
			if val - ALetter[key] >= 0 {
				cnt += val - ALetter[key]
			} else {
				cnt += ALetter[key] - val
			}
			ATrue[key] = false
			BTrue[key] = false
		} else if BTrue[key] {
			cnt += val
			BTrue[key] = false
		}
	}

	fmt.Println(cnt)
}