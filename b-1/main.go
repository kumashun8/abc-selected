// 黒板に NN 個の正の整数 A1,…,ANA1,…,AN が書かれています。
// すぬけ君は，黒板に書かれている整数がすべて偶数であるとき，次の操作を行うことができます。
// - 黒板に書かれている整数すべてを，2 で割ったものに置き換える。
// すぬけ君は最大で何回操作を行うことができるかを求めてください。

package main

import "fmt"

func shiftOnly(nums []int, length int, score int) int {
	var i int
	for i = 0; i < length; i++ {
		if nums[i]%2 != 0 {
			return score
		}
		nums[i] = nums[i] / 2
	}

	return shiftOnly(nums, length, score+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(shiftOnly(nums, n, 0))
}
