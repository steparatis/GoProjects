package jointwoslice

import "fmt"

func readSlice(s []int, ch chan int, result *[]int) {
	for _,v := range s {
		ch <- v
	}
	close(ch)
	return
}

func Join() {
	odd := []int{1,3,5,7,9}
	even := []int{2,4,6,8,10}
	chOdd := make(chan int)
	chEven := make(chan int)
	var result []int

	go readSlice(odd, chOdd, &result)
	go readSlice(even, chEven, &result)

	for {
		vOdd, okOdd := <-chOdd		
		vEven, okEven := <-chEven
		if okOdd {
			result = append(result, vOdd)
		}
		if okEven {
			result = append(result, vEven)
		}
		if !okEven && !okOdd {
			break
		}
	}	
	fmt.Println(result)
	
}