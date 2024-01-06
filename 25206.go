package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var subject, grade string 	
	var credit, sum_grade, sum_credit float64
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<20;i++ {
		fmt.Fscanln(reader, &subject, &credit, &grade)

			switch grade {
			case "A+" :
				sum_credit += credit
				sum_grade += credit * 4.5
			case "A0" :
				sum_credit += credit
				sum_grade += credit * 4
			case "B+" :
				sum_credit += credit
				sum_grade += credit * 3.5
			case "B0" :
				sum_credit += credit
				sum_grade += credit * 3
			case "C+" :
				sum_credit += credit
				sum_grade += credit * 2.5
			case "C0" :
				sum_credit += credit
				sum_grade += credit * 2
			case "D+" :
				sum_credit += credit
				sum_grade += credit * 1.5
			case "D0" :
				sum_credit += credit
				sum_grade += credit * 1
			case "F" :
				sum_credit += credit
				sum_grade += credit * 0
			}			
		fmt.Println(sum_grade, sum_credit)
	}
	fmt.Println(sum_grade/sum_credit)
}