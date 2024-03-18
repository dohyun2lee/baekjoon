package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var N, New, P, x, cnt, rank int
	fmt.Fscan(r, &N, &New, &P)
  for i := 0; i < N; i++ {
    fmt.Fscan(r, &x)
    if x >= New {
      cnt++
    }
    if x > New {
      rank++
    }
  }
  if cnt >= P {
    fmt.Print(-1)
  } else {
    fmt.Print(rank+1)
  }
}
