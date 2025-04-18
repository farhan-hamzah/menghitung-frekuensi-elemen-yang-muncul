package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, j, jum int
	jum = 1
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		sudah := false
		for j = 0; j < i; j++ {
			if A[i] == A[j] {
				sudah = true
				break
			}
		}
		if sudah {
			continue
		}

		for j = i+1; j < n; j++{
			if A[i] == A[j]{
				jum++
			}
		}
		fmt.Println(A[i], " Muncul ", jum)
		jum = 1
	}
}