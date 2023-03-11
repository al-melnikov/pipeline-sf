package main

import (
	"fmt"
	"pipeline/pipeline"
)

func consume(done <-chan bool, source <-chan int) {
	for {
		select {
		case data := <-source:
			fmt.Printf("Обработаны данные: %d\n", data)
		case <-done:
			return
		}
	}
}

func main() {
	var ns pipeline.NegativeFilter
	var sf pipeline.SpecialFilter
	var bs pipeline.BufferStage

	source, done := pipeline.ReceiveData()

	pl := pipeline.New(done, &ns, &sf, &bs)

	pipeline.BufferSize = 5          //размер буфера 5
	pipeline.BufferDrainInterval = 5 // интервал в секундах между просмотрами буфера
	consume(done, pl.Run(source))
}
