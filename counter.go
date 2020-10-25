package counter

type Counter struct {
	value int
	queue chan int
	stop  chan bool
}

func NewCounter(initValue int, chanSize int) *Counter {
	queue := make(chan int, chanSize)
	stop := make(chan bool)
	counter := &Counter{
		value: initValue,
		queue: queue,
		stop:  stop,
	}
	go counter.Resume()
	return counter
}

func (counter *Counter) Stop() {
	counter.stop <- false
}

func (counter *Counter) Resume() {
	for {
		select {
		case <-counter.stop:
			{
				return
			}
		default:
			{
				select {
				case d := <-counter.queue:
					{
						counter.value += d
						break
					}
				case <-counter.stop:
					{
						return
					}
				}
			}
		}
	}
}

func (counter *Counter) Value() int {
	return counter.value
}

func (counter *Counter) Flush() {
	counter.Stop()
	for {
		select {
		case d := <-counter.queue:
			{
				counter.value += d
				break
			}
		default:
			{
				go counter.Resume()
				return
			}
		}
	}
}

func (counter *Counter) Plus(d int) {
	counter.queue <- d
}

func (counter *Counter) Dispose() {
	counter.Stop()
}
