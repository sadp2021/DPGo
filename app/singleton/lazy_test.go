package singleton_test

import (
	"testing"

	"github.com/sadp2021/DPGo/app/singleton"
	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, singleton.GetLazyInstance(), singleton.GetLazyInstance())
		}
	})
}
