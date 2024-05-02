package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var Name string
	var Age, Weight int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscanln(reader, &Name, &Age, &Weight)

		if Name == "#" {
			break
		}

		if Age > 17 || Weight >= 80 {
			fmt.Fprintln(writer, Name, "Senior")
		} else {
			fmt.Fprintln(writer, Name, "Junior")
		}
	}
}