package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)


func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &s)

	fmt.Println(reflect.TypeOf(s[0]), s[1])
}
