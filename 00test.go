package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var pool = []string{"a", "b", "c", "a", "d", "c"}

	deleteDuplicateItem(pool)

	fmt.Println(pool)
}

func deleteDuplicateItem(arr []string) []string {
    ret := []string{}
    m := make(map[string]int)
    
    for _, val := range arr {
    	if _, ok := m[val]; !ok {
			fmt.Println("val:", val)
			fmt.Println("m[val]:", m[val])
            m[val] = 1
            ret = append(ret, val)
    	}
    }

	fmt.Println(ret)
    
    return ret
}