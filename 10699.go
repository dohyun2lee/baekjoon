package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ans := now.Format("2006-01-02")

	fmt.Println(ans)
}
