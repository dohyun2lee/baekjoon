package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fscanln(reader, &T)

		if T == 0 {
			break
		}

		var ans int

		for j:=1;j<=T;j++ {
			ans += j
		}

		fmt.Println(ans)
	}
}