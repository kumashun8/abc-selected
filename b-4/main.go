// 2 人が交互に N 枚のデッキから 1 枚ずつカードを取っていきます。
// 2 人がすべてのカードを取ったときゲームは終了し、取ったカードの数の合計がその人の得点になります。
// 2 人とも自分の得点を最大化するように最適戦略をとったとき、Alice は Bob より何点多くの得点を獲得できるかを求めてください。

package main

import (
	"fmt"
	"sort"
)

func cardGameForTwo(n int, deck []int) int {
	scores := []int{0, 0}
	sort.Slice(deck, func(i int, j int) bool { return deck[i] > deck[j] })

	for i := 0; i < n; i++ {
		scores[i%2] += deck[i]
	}

	return scores[0] - scores[1]
}

func main() {
	var n int
	fmt.Scan(&n)

	deck := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&deck[i])
	}

	fmt.Println(cardGameForTwo(n, deck))
}
