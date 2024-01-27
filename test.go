package main

import (
	// "container/list"
	"fmt"
	"sort"
)

type Score struct {
	Cluster string
	Score   float32
}

func (s ScoreTable) Len() int {
	return len(s)
}

func (s ScoreTable) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s ScoreTable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ScoreTable []Score

var score_table = make(map[string]float32)

func main() {
	score_table["a-cluster"] = 30.0
	score_table["b-cluster"] = 40.0
	score_table["c-cluster"] = 20.0
	score_table["d-cluster"] = 50.0

	score_sort := SortScore()
	for _, score := range score_sort {
		fmt.Printf("%s %2.2f\n", score.Cluster, score.Score)
	}
}

func SortScore() ScoreTable {

	// 1. 구조체 배열 만들기
	sorted := make(ScoreTable, len(score_table))

	// 2. map에 있는 값을 구조채 배열에 입력하기
	for cluster, score := range score_table {
		sorted = append(sorted, Score{cluster, score})
	}

	// 3. sort 패키지의 Sort함수를 이용해 정렬하기
	sort.Sort(sorted)

	return sorted
}
