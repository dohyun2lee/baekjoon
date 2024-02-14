package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type word struct {
	spell string
	count int
}

func main() {
	var S, ans, chk string
	var Ans []string
	var tf bool
	var cnt int
	var wordTable []word
	var letter map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)
	s := strings.Split(S, "")

	for i := 0; i < len(s); i++ {
		letter[s[i]]++ // 알파벳 입력 받은 개수 넣기
	}

	for k, v := range letter {
		letter[k] = v / 2
		if letter[k] == 0 {
			cnt++
		}
		if cnt >= 2 {
			tf = true // 홀수 입력이 2개 이상체크
			break
		}
		wordTable = append(wordTable, word{k, v})
	}

	sort.Slice(wordTable, func(i, j int) bool { //wordTable 정렬 / count가 높은 순, 같으면 알파벳 빠른순부터
		if wordTable[i].count/2 == wordTable[j].count/2 { // /2한 이유는 홀수까지 같이 비교하기 위함
			return wordTable[i].spell < wordTable[j].spell
		}
		return wordTable[i].count/2 > wordTable[j].count/2
	})

	if tf {
		fmt.Println("I'm Sorry Hansoo") // 홀수가 2개 이상이면 I'm Sorry Hansoo 출력
	} else {
		for i := 0; i < len(wordTable); i++ {
			if wordTable[i].count%2 == 0 { //wordTable의 요소가 짝수면 Ans에 알파벳 개수의 반 만큼 추가
				for j := 0; j < wordTable[i].count/2; j++ {
					Ans = append(Ans, wordTable[i].spell) 
				}
			} else { //wordTable의 요소가 홀수면 Ans에 알파벳 개수의 반 만큼 추가한 뒤, chk에 홀수요소 할당
				for j := 0; j < wordTable[i].count/2; j++ {
					Ans = append(Ans, wordTable[i].spell)
				}
				chk = wordTable[i].spell
			}
		}

		for i := 0; i < len(Ans); i++ { //ans에 Ans요소들 다 추가후, chk추가한 뒤 Ans를 뒤집어서 입력
			ans += Ans[i]
		}
		ans += chk
		for i := len(Ans) - 1; i >= 0; i-- {
			ans += Ans[i]
		}

		fmt.Println(ans)
	}
}
