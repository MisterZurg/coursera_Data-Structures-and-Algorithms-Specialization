package Week_1_Programming_Challenges

import (
	"fmt"
	"math/big"
)

func MaximumPairwiseProduct() {
	var n int
	fmt.Scan(&n)

	var inputDataSet = make([]int64, n, n)

	// Заполним слайс значениями
	for i := 0; i < n; i++ {
		fmt.Scan(&inputDataSet[i])
	}

	var imax1, imax2, imin1, imin2 int = -1, -1, -1, -1
	var firstCheck = true
	for idx, num := range inputDataSet {
		if firstCheck {
			firstCheck = false
			imax1 = idx
			imin1 = idx
			continue
		}
		if num > inputDataSet[imax1] {
			imax2 = imax1
			imax1 = idx
			// Первым условием мы обходим несуществующий индекс
		} else if imax2 == -1 || num >= inputDataSet[imax2] {
			imax2 = idx
		}
		// Min
		if num < inputDataSet[imin1] {
			imin2 = imin1
			imin1 = idx
			// Первым условием мы обходим несуществующий индекс
		} else if imin2 == -1 || num <= inputDataSet[imin2] {
			imin2 = idx
		}
	}

	if inputDataSet[imin2]*inputDataSet[imin1] > inputDataSet[imax2]*inputDataSet[imax1] {
		// Произведение минимальных
		fmt.Println(big.NewInt(int64(inputDataSet[imax2]) * int64(inputDataSet[imax1])))
	} else {
		// Произведение максимальных
		fmt.Println(big.NewInt(int64(inputDataSet[imax2]) * int64(inputDataSet[imax1])))
	}
}
