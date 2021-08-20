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

func Map(done <-chan interface{}, inputChannel <-chan interface{}, mapFunc func(interface{}) interface{}) <-chan interface{} {

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			case outputChannel <- mapFunc(value):
			}
		}
	}()
	return outputChannel
}

func Take(done <-chan interface{}, inputChannel <-chan interface{}, nElements int) <-chan interface{} {

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for count := 1; count <= nElements; count++ {
			select {
			case <-done:
				return
			case outputChannel <- <-inputChannel:
			}
		}
	}()
	return outputChannel
}

func DropAll(done <-chan interface{}, inputChannel <-chan interface{}, element int) <-chan interface{} {

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				if value != element {
					outputChannel <- value
				}
			}
		}
	}()
	return outputChannel
}

func Merge(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		publishFrom := func(ch <-chan interface{}) {
			for value := range ch {
				select {
				case <-done:
					return
				case outputChannel <- value:
				}
			}
		}
		for _, channel := range channels {
			publishFrom(channel)
		}
	}()
	return outputChannel
}
