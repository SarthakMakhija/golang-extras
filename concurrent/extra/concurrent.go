package extra

func Repeat(done <-chan interface{}, fn func() interface{}) <-chan interface{} {

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for {
			select {
			case <-done:
				return
			case outputChannel <- fn():
			}
		}
	}()
	return outputChannel
}

func Map(done <-chan interface{}, incoming <-chan interface{}, mapFunc func(interface{}) interface{}) <-chan interface{} {

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range incoming {
			select {
			case <-done:
				return
			case outputChannel <- mapFunc(value):
			}
		}
	}()
	return outputChannel
}
