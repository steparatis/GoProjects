package generatorPseudoRandomNumbers

import (
	"fmt"
	"math/rand"
	"time"

	joinchanels "github.com/adminsemy/golangTests/joinChanels"
)

//Получаем число, до которого надо взять случайное
//число и возваращаем псевдо-случайное число
func Generating(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func writeRandomNumber(ch chan<- int, number int, repeat int) {
	for i := 0; i <= repeat; i++ {
		n := Generating(number)
		ch <- n
	}
	close(ch)
}

func Start() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	
	//join := joinchanels.JoinChanels(a, b, c)
	go writeRandomNumber(a, 100, 2)
	go writeRandomNumber(b, 10, 2)
	go writeRandomNumber(c, 100, 2)
	joinCannel := joinchanels.JoinChanels(a, b, c)

	for v:= range joinCannel {
		fmt.Println(v)
	}

}