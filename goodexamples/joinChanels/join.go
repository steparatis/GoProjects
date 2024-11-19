package joinchanels

import "sync"

func JoinChanels(chs ...<-chan int) chan int {
	merge := make(chan int)
	go func (merge chan int)  {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))
	
		for _, v := range chs {
			go func (ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					merge <-id
				}			
			}(v, wg)
		}
		wg.Wait()
		close(merge)	
	}(merge)
	
	return merge
}