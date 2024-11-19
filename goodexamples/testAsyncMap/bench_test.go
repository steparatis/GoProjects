package testAsyncMap_test

import (
	"testing"

	"github.com/adminsemy/golangTests/testAsyncMap"
)

func BenchmarkXxx(b *testing.B) {
	async := testAsyncMap.NewAsyncMap()
	for i := 0; i < b.N; i++ {
		emulateLoad(&testing.T{}, async)
	}
}
