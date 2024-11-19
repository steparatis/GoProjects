//Алгоритм к задаче ханойские башни.
package hanoi

import "fmt"

func Hanoi(countDisks, os1, os2 int) {
	if countDisks == 1 {
		fmt.Printf("Переложи диск 1 с %v на %v\n", os1, os2)
		return
	}
	tmp := 6 - os1 - os2
	Hanoi(countDisks-1, os1, tmp)
	fmt.Printf("Переложи диск %v с %v на %v\n", countDisks, os1, os2)
	Hanoi(countDisks-1, tmp, os2)

}
