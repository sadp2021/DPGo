package singleton_test

import (
	"testing"

	"github.com/sadp2021/DPGo/app/singleton"
	"github.com/stretchr/testify/assert"
)

func TestGetHungryInstance(t *testing.T) {
	assert.Equal(t, singleton.GetHungryInstance(), singleton.GetHungryInstance())
}

func BenchmarkGetHungryInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, singleton.GetHungryInstance(), singleton.GetHungryInstance())
		}
	})
}
