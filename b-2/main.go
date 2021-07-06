// 500 円玉を AA 枚、100 円玉を BB 枚、50 円玉を CC 枚持っています。
// これらの硬貨の中から何枚かを選び、合計金額をちょうど XX 円にする方法は何通りあるでしょうか？

package main

import "fmt"

func findCoinPattern(a int, b int, c int, total int) int {
	var score = 0

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if canMakeTotal(i, j, k, total) {
					score++
				}
			}
		}
	}

	return score
}

func canMakeTotal(a int, b int, c int, total int) bool {
	return (500*a + 100*b + 50*c) == total
}

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	fmt.Println(findCoinPattern(a, b, c, x))
}
