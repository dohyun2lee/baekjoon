if i == 0 {
	for j:=0;j<2;j++ {
		if start[j] == "0" {
			start[j] = "1"
		} else {
			start[j] = "0"
		}
	}
} else if i == N-1 {
	for j:=N-2;j<N;j++ {
		if start[j] == "0" {
			start[j] = "1"
		} else {
			start[j] = "0"
		}
	}
} else {
	for j:=i-1;j<=i+1;j++ {
		if start[j] == "0" {
			start[j] = "1"
		} else {
			start[j] = "0"
		}
	}
}