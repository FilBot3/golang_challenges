package sumOfNumbersToN

import (
	"fmt"
)

func SumOfNumbersToN(n int64) (total int64) {
	total = 0
	var i int64
	for i = 0; i <= n; i++ {
		total = total + i
	}
	fmt.Println(total)

	return total
}
