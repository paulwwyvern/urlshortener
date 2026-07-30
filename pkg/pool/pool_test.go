package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	id   int
	data int
}

func (p *testStruct) Reset() {
	p.data = 0
}

func TestPool(t *testing.T) {
	id := 0
	pool := New(func() *testStruct {
		iid := id
		id++
		return &testStruct{
			id: iid,
		}
	})

	fi := pool.Get()
	se := pool.Get()

	assert.Equal(t, 0, fi.id)
	assert.Equal(t, 1, se.id)

	fi.data = 42

	pool.Put(fi)
	pool.Put(se)

	foo := pool.Get()
	bar := pool.Get()
	baz := pool.Get()

	assert.Equal(t, 1, foo.id)
	assert.Equal(t, 0, foo.data)

	assert.Equal(t, 0, bar.id)
	assert.Equal(t, 0, bar.data)

	assert.Equal(t, 2, baz.id)
	assert.Equal(t, 0, baz.data)
}

func TestPool_NilFunc(t *testing.T) {
	pool := New[*testStruct](nil)

	foo := &testStruct{
		id:   1,
		data: 42,
	}

	pool.Put(foo)

	bar := pool.Get()
	assert.Equal(t, 1, bar.id)
	assert.Equal(t, 0, bar.data)

	assert.Nil(t, pool.Get())

}
