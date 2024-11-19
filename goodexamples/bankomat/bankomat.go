/*
Банкомат, который выдает минимальную возможную комбинацию денег определенной суммы
*/

package bankomat

import "fmt"

const GOOGOL int32 = 99999999

var cashe = make(map[int32]int32)

func GiveMeMoney(sum int32, banknots []int32) int32 {
	casheSum, ok := cashe[sum]
	if ok {
		return casheSum
	}
	if sum < 0 {
		return GOOGOL
	}
	if sum == 0 {
		return 0
	}
	var min int32 = GOOGOL
	for i := 0; i < len(banknots); i++ {
		nextSum := sum - banknots[i]
		result := GiveMeMoney(nextSum, banknots)
		fmt.Printf("Текущая сумма - %v, следующая сумма - %v, результат - %v\n", sum, nextSum, result)
		if result < min {
			min = result
		}
	}
	cashe[sum] = min + 1
	fmt.Println("Сумма на выход - ", cashe[sum])
	return cashe[sum]
}
