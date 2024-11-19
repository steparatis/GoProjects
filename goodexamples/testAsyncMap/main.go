package testAsyncMap

import (
	"fmt"
)

func main() {
	async := make(map[string]string)
	async["one"] = "one"
	for i := 0; i <= 100; i++ {
		go func() {
			res := async["one"]
			fmt.Println(res)
		}()
	}
}
