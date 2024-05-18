package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var S string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var vowelCnt int
		fmt.Fscanln(reader, &S)
		if S == "end" {
			break
		} 

		for i:=0;i<len(S);i++ {
			
		}
	}
}