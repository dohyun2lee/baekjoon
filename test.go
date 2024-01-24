package main

import (
	//	"bufio"
	"os"
	"bufio"
	"fmt"
	//"reflect"
	// 	"strconv"
	//	"sort"
	"strings"
)

func main() {
	// var X []int
	// var a int
	// reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	chk := scanner.Text()
	var arr []string = strings.Split(chk, " ")

	if arr[0] == arr[1] {
		fmt.Println(1)
	}
}
