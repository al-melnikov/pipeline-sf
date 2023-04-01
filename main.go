package main

import (
	"fmt"
	"log"
	"pipeline/pipeline"
)

func consume(done <-chan bool, source <-chan int) {
	for {
		select {
		case data := <-source:
			log.Printf("Обработаны данные: %d\n", data)
			fmt.Printf("Обработаны данные: %d\n", data)
		case <-done:
			log.Printf("consume finished\n")
			return
		}
	}
}

func main() {
	log.Println("program started")
	var ns pipeline.NegativeFilter
	var sf pipeline.SpecialFilter
	var bs pipeline.BufferStage

	source, done := pipeline.ReceiveData()

	pl := pipeline.New(done, &ns, &sf, &bs)

	pipeline.BufferSize = 5          //размер буфера 5
	pipeline.BufferDrainInterval = 5 // интервал в секундах между просмотрами буфера
	consume(done, pl.Run(source))
	log.Println("program finished")
}
