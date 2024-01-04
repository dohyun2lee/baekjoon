package main 

import (
	"fmt"
	"bufio"
	"os"
	//"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}