	package main

	import (
		"bufio"
		"os"
		"fmt"
		"sort"
	)

	type person struct {
		age, timing int
		name string
	}

	func main() {
		var N, age int
		var name string
		reader := bufio.NewReader(os.Stdin)
		writer := bufio.NewWriter(os.Stdout)
		defer writer.Flush()

		fmt.Fscanln(reader, &N)
		var member []person = make([]person, N)

		for i:=0;i<N;i++ {
			fmt.Fscanln(reader, &age, &name)
			member[i].age = age
			member[i].name = name
			member[i].timing = i
		}

		sort.Slice(member, func(i, j int) bool {
			if member[i].age == member[j].age {
				return member[i].timing < member[j].timing
			}
			return member[i].age < member[j].age
		})

		for i:=0;i<N;i++ {
			fmt.Fprintln(writer, member[i].age, member[i].name)
		}
	}