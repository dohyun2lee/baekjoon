package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cnt, tmp, chkTmp int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	chk := scanner.Text()

	for i := 0; i <= len(s)-len(chk); i++ {
		tmp = 0
		chkTmp = 0
		// fmt.Println("i:",i)
		for j := i; j < i+len(chk); j++ {
			// fmt.Println("s:",string(s[j]))
			// fmt.Println("cnk:",string(chk[chkTmp]))
			if s[j] == chk[chkTmp] {
				tmp++
			} else {
				break
			}
			chkTmp++
		}
		if tmp == len(chk) {
			cnt++
			i += len(chk) - 1
		}
	}

	fmt.Println(cnt)
}
