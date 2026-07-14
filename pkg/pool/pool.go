package pool

import "sync"

type Resetter interface {
	Reset()
}

type Pool[R Resetter] struct {
	mu      sync.Mutex
	objects []R
	newFn   func() R
}

func New[R Resetter](newFn func() R) *Pool[R] {
	return &Pool[R]{newFn: newFn}
}

func (p *Pool[R]) Get() R {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.objects) == 0 {
		return p.newFn()
	}

	last := len(p.objects) - 1
	obj := p.objects[last]
	p.objects = p.objects[:last]
	return obj
}

func (p *Pool[R]) Put(obj R) {
	obj.Reset()

	p.mu.Lock()
	p.objects = append(p.objects, obj)
	p.mu.Unlock()
}
