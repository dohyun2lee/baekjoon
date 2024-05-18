package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int 
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i:=1;i<=N;i++ {
		scanner.Scan()
		S := scanner.Text()
		fmt.Fprintf(writer, "%d. %s\n", i, S)
	}
}