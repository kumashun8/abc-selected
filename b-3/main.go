// 1  以上 N 以下の整数のうち、10 進法で各桁の和が A 以上 B 以下であるものについて、総和を求めてください。

package main

import "fmt"

func someSums(a int, b int, n int) int {
	var sum = 0

	for i := a; i <= n; i++ {
		if c := calcDigitSum(i); a <= c && c <= b {
			sum += i
		}
	}

	return sum
}

func calcDigitSum(n int) int {
	var sum = 0

	for i := n; i >= 1; i = i / 10 {
		sum += i % 10
	}

	return sum
}

func main() {
	var a, b, n int
	fmt.Scan(&n, &a, &b)

	fmt.Println(someSums(a, b, n))
}
