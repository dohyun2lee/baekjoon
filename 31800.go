package main 

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, first, second, x, chance, ans int
	var profit []int
	var price []int
	var profit_2 []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		if i != N-1 {
			fmt.Fscan(reader, &x)
		} else {
			fmt.Fscanln(reader, &x)
		}

		profit = append(profit, x)
		profit_2 = append(profit_2, x)
	}

	for i:=0;i<N;i++ {
		if i != N-1 {
			fmt.Fscan(reader, &x)
		} else {
			fmt.Fscanln(reader, &x)
		}

		price = append(price, x)
	}

	sort.Ints(profit_2)

	first = profit_2[len(profit_2)-1]
	second = profit_2[len(profit_2)-2]

	for i:=0;i<N;i++ {
		if profit[i] == first {
			chance = second - price[i]
		} else {
			chance = first - price[i]
		}

		ans = profit[i] - chance - price[i]

		fmt.Fprintf(writer, "%d ", ans)
	}
}