package testAsyncMap_test

import (
	"testing"

	"github.com/adminsemy/golangTests/testAsyncMap"
)

func TestAsyncMap(t *testing.T) {
	t.Run("parallel test asycn Map", func(t *testing.T) {
		async := testAsyncMap.NewAsyncMap()
		emulateLoad(t, async)
	})
}
