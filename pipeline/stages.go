// Здесь описаны стадии пайплайна, они имплиментируют интерфейс Stage

package pipeline

import (
	"pipeline/ringbuffer"
	"time"
)

var BufferSize int = 10
var BufferDrainInterval int = 30 // в секундах

type Stage interface {
	run(<-chan bool, <-chan int) <-chan int
}

type NegativeFilter struct{}
type SpecialFilter struct{}
type BufferStage struct{}

func (nf *NegativeFilter) run(done <-chan bool, source <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for {
			select {
			case data := <-source:
				if data > 0 {
					select {
					case res <- data:
					case <-done:
						return
					}
				}
			case <-done:
				return
			}
		}
	}()
	return res
}

func (sf *SpecialFilter) run(done <-chan bool, source <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for {
			select {
			case data := <-source:
				if data != 0 && data%3 == 0 {
					select {
					case res <- data:
					case <-done:
						return
					}
				}
			case <-done:
				return
			}
		}
	}()
	return res
}

func (bf *BufferStage) run(done <-chan bool, source <-chan int) <-chan int {
	res := make(chan int)
	buffer := ringbuffer.NewRingBuffer(BufferSize)
	go func() {
		for {
			select {
			case data := <-source:
				buffer.Push(data)
			case <-done:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Duration(BufferDrainInterval) * time.Second):
				bufferData := buffer.Get()
				for _, data := range bufferData {
					select {
					case res <- data:
					case <-done:
						return
					}
				}
			case <-done:
				return
			}
		}
	}()
	return res
}
