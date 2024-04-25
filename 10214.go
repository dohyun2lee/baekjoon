package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, korea, yonsei, koreaCnt, yonseiCnt int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		koreaCnt = 0
		yonseiCnt = 0
		for j:=0;j<9;j++ {
			fmt.Fscanln(reader, &yonsei, &korea)
			koreaCnt += korea
			yonseiCnt += yonsei
		}

		if koreaCnt > yonseiCnt {
			fmt.Fprintln(writer, "Korea")
		} else if yonseiCnt > koreaCnt {
			fmt.Fprintln(writer, "Yonsei")
		} else {
			fmt.Fprintln(writer, "Draw")
		}
	}
}
