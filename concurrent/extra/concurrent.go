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
