// Здесь описан кольцевой буффер
package ringbuffer

import (
	"container/ring"
	"sync"
)

type RingBuffer struct {
	r *ring.Ring
	m sync.Mutex
}

func NewRingBuffer(size int) *RingBuffer {
	ringBuf := RingBuffer{r: ring.New(size), m: sync.Mutex{}}
	return &ringBuf
}

func (r *RingBuffer) Push(el int) {
	r.m.Lock()
	defer r.m.Unlock()
	r.r.Value = el
	r.r = r.r.Next()
}

func (r *RingBuffer) Get() []int {
	r.m.Lock()
	defer r.m.Unlock()
	n := r.r.Len()
	res := make([]int, 0, n)
	for j := 0; j < n; j++ {
		v := r.r.Value
		r.r.Value = nil
		if v != nil {
			res = append(res, v.(int))
		}
		r.r = r.r.Next()
	}
	return res
}
