// 二つの正整数 aa, bb が与えられます。
// aa と bb の積が偶数か奇数か判定してください。

package main

import "fmt"

func product(aa int, bb int) string {
	prd := aa * bb

	if prd%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	var aa, bb int
	fmt.Scan(&aa, &bb)

	fmt.Println(product(aa, bb))
}
