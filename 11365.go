package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var ans string
		scanner.Scan()
		S := scanner.Text()

		if S == "END" {
			break
		}

		for i:=len(S)-1;i>=0;i-- {
			ans += string(S[i])
		}

		fmt.Fprintln(writer, ans)
	}
}