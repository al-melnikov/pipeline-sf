// Здесь описан пайплайн в целом и методы для работы с ним

package pipeline

type Pipeline struct {
	stages []Stage
	done   <-chan bool
}

// Создает новый пайплайн и возвращает указатель на него
func New(done <-chan bool, stages ...Stage) *Pipeline {
	return &Pipeline{stages: stages, done: done}
}

// Запускает стадии пайплайна поочередно
func (p *Pipeline) Run(source <-chan int) <-chan int {
	var c <-chan int = source
	for index := range p.stages {
		c = p.runStage(p.stages[index], c)
	}
	return c
}

func (p *Pipeline) runStage(stage Stage, sourceChan <-chan int) <-chan int {
	// запустть канал и вернуть его результат
	return stage.run(p.done, sourceChan)
}
