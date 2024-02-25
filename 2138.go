package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N, cnt1, cnt2, ans int
	var s1, s2 string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	fmt.Fscanln(reader, &s1)
	fmt.Fscanln(reader, &s2)

	if s1 == s2 {
		ans = 0
	} else {
		start := strings.Split(s1, "")
		end := strings.Split(s2, "")

		if N == 2 {
			for i := 0; i < N; i++ {
				if start[i] == "0" {
					start[i] = "1"
				} else {
					start[i] = "0"
				}
			}
		} else {
			// 0번 안누름
			if start[0] == end[0] {
				for i := 2; i < N; i++ {
					if i != N-1 {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i+1; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					} else {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					}
					
					cnt1++
					if strings.Join(start, "") == s2 {
						break
					}
				}
			} else {
				for i := 1; i < N; i++ {
					if i != N-1 {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i+1; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					} else {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					}
					
					cnt1++
					if strings.Join(start, "") == s2 {
						break
					}
				}
			}

			// 0번 누름
			cnt2++
			for i := 0; i < 2; i++ {
				if start[i] == "0" {
					start[i] = "1"
				} else {
					start[i] = "0"
				}
			}

			if start[0] == end[0] {
				for i := 2; i < N; i++ {
					if i != N-1 {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i+1; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					} else {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					}
					
					cnt2++
					if strings.Join(start, "") == s2 {
						break
					}
				}
			} else {
				for i := 1; i < N; i++ {
					if i != N-1 {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i+1; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					} else {
						if start[i-1] != end[i-1] {
							for j := i - 1; j <= i; j++ {
								if start[j] == "0" {
									start[j] = "1"
								} else {
									start[j] = "0"
								}
							}
						}
					}
					
					cnt2++
					if strings.Join(start, "") == s2 {
						break
					}
				}
			}

			if cnt1 > cnt2 {
				ans = cnt2
			} else {
				ans = cnt1
			}
		}
	}

	fmt.Println(ans)
}
