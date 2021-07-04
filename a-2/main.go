// 0 と 1 のみから成る 3 桁の番号 s が与えられます。1 が何個含まれるかを求めてください。
package main

import (
	"fmt"
	"strings"
)

func placingMerbles(s string) int {
	return strings.Count(s, "1")
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(placingMerbles(s))
}
