package testAsyncMap_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/adminsemy/golangTests/testAsyncMap"
	"github.com/stretchr/testify/assert"
)

func emulateLoad(t *testing.T, c testAsyncMap.CasheMap) {
	wg := &sync.WaitGroup{}
	for i := 0; i <= 1_0; i++ {
		key := fmt.Sprintf("#%v-key", i)
		value := fmt.Sprintf("#%v-value", i)
		wg.Add(1)
		go func(k, value string) {
			err := c.Write(k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key, value)

		for y := 0; y <= 10_00; y++ {
			wg.Add(1)
			go func(k, v string) {
				val, err := c.Read(k)
				if err == nil {
					assert.Equal(t, v, val)
				}
				wg.Done()
			}(key, value)
		}

		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)
	}
	wg.Wait()
}
