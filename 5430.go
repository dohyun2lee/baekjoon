package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var N int
		var cmd string
		var err bool
		fmt.Fscanln(reader, &cmd)
		fmt.Fscanln(reader, &N)
		scanner.Scan()

		command := strings.Split(cmd, "")
		x := scanner.Text()
		X := strings.Split(x, ",")

		X = removeRest(X)

		for j := 0; j < len(command); j++ {
			if command[j] == "R" {
				X = R(X)
			} else {
				if len(X) == 0 {
					err = true
					break
				} else {
					X = X[1:]
				}
			}
		}

		if err {
			fmt.Fprintln(writer, "error")
		} else {
			fmt.Fprintf(writer, "[")
			for j := 0; j < len(X); j++ {
				if j != len(X)-1 {
					fmt.Fprintf(writer, "%s,", X[j])
				} else {
					fmt.Fprintf(writer, "%s", X[j])
				}
			}
			fmt.Fprintf(writer, "]\n")
		}
	}
}

func removeRest(X []string) (res []string) {
	if len(X) == 1 {
		var tmp []string
		a := strings.Split(X[0], "")
		for i := 0; i < len(a); i++ {
			if a[i] != "[" && a[i] != "]" {
				tmp = append(tmp, a[i])
			}
		}
		if len(tmp) == 0 {
			res = nil
		} else {
			res = append(res, strings.Join(tmp, ""))
		}
	} else if len(X) == 2 {
		first := strings.Split(X[0], "")
		last := strings.Split(X[len(X)-1], "")

		first = first[1:]
		last = last[:len(last)-1]

		res = append(res, strings.Join(first, ""))
		res = append(res, strings.Join(last, ""))
	} else {
		first := strings.Split(X[0], "")
		last := strings.Split(X[len(X)-1], "")

		first = first[1:]
		last = last[:len(last)-1]

		X = X[1:]
		X = X[:len(X)-1]

		res = append(res, strings.Join(first, ""))
		res = append(res, X...)
		res = append(res, strings.Join(last, ""))
	}

	return res
}

func R(X []string) (res []string) {
	for i := len(X) - 1; i >= 0; i-- {
		res = append(res, X[i])
	}

	return res
}

// func D(X []string) (res []string, err bool) {
// 	if len(X) != 0 {
// 		res = X[1:]
// 		err = false
// 	} else {
// 		res = nil
// 		err = true
// 	}

// 	return res, err
// }
