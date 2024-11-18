package main

import "fmt"

func main() {
	var s = "III"
	fmt.Println(RimConv(s))
}
func RimConv(s string) int {
	list := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var ans = list[string(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		if list[string(s[i])] < list[string(s[i+1])] {
			ans -= list[string(s[i])]
		} else {
			ans += list[string(s[i])]
		}
	}
	return ans
}
