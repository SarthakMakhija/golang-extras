package extra

import (
	"fmt"
	"gitlab.com/stone.code/assert"
)

func Repeat(
	done <-chan interface{},
	generatorFunction func() interface{},
) <-chan interface{} {

	assertNonNilChannels(done)

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for {
			select {
			case <-done:
				return
			case outputChannel <- generatorFunction():
			}
		}
	}()
	return outputChannel
}

func Map(
	done <-chan interface{},
	inputChannel <-chan interface{},
	mapFunction func(interface{}) interface{},
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			case outputChannel <- mapFunction(value):
			}
		}
	}()
	return outputChannel
}

func Filter(
	done <-chan interface{},
	inputChannel <-chan interface{},
	filterFunction func(interface{}) bool,
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				if filterFunction(value) {
					outputChannel <- value
				}
			}
		}
	}()
	return outputChannel
}

func RunningReduce(
	done <-chan interface{},
	inputChannel <-chan interface{},
	initialValue interface{},
	aggregateFunction func(aggregate interface{}, value interface{}) interface{},
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

	outputChannel := make(chan interface{})
	go func() {
		aggregate := initialValue
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				aggregate = aggregateFunction(aggregate, value)
				outputChannel <- aggregate
			}
		}
	}()
	return outputChannel
}

func Skip(
	done <-chan interface{},
	inputChannel <-chan interface{},
	skipFunction func(interface{}) bool,
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				if !skipFunction(value) {
					outputChannel <- value
				}
			}
		}
	}()
	return outputChannel
}

func Take(
	done <-chan interface{},
	inputChannel <-chan interface{},
	nElements int,
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

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

func TakeWhile(
	done <-chan interface{},
	inputChannel <-chan interface{},
	condition func(interface{}) bool,
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				if condition(value) {
					outputChannel <- value
				} else {
					return
				}
			}
		}
	}()
	return outputChannel
}

func DropAll(
	done <-chan interface{},
	inputChannel <-chan interface{},
	element int,
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

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

func Reverse(
	done <-chan interface{},
	inputChannel <-chan interface{},
) <-chan interface{} {

	assertNonNilChannels(done, inputChannel)

	outputChannel := make(chan interface{})
	go func() {
		defer close(outputChannel)
		var elements []interface{}

		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				elements = append(elements, value)
			}
		}
		for index := len(elements) - 1; index >= 0; index-- {
			select {
			case <-done:
				return
			case outputChannel <- elements[index]:
			}
		}
	}()
	return outputChannel
}

func Merge(
	done <-chan interface{},
	channels ...<-chan interface{},
) <-chan interface{} {

	assertNonNilChannels(done)
	assertNonNilChannels(channels...)

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

func Tee(
	done <-chan interface{},
	inputChannel <-chan interface{},
) (<-chan interface{}, chan interface{}) {

	assertNonNilChannels(done, inputChannel)

	outputChannel1 := make(chan interface{})
	outputChannel2 := make(chan interface{})

	go func() {
		defer close(outputChannel1)
		defer close(outputChannel2)

		for value := range inputChannel {
			select {
			case <-done:
				return
			default:
				for count := 1; count <= 2; count++ {
					select {
					case <-done:
						return
					case outputChannel1 <- value:
					case outputChannel2 <- value:
					}
				}
			}
		}
	}()
	return outputChannel1, outputChannel2
}

func assertNonNilChannels(channels ...<-chan interface{}) {
	for _, ch := range channels {
		assert.Assert(ch != nil, fmt.Sprintf("Channel %v must not be nil", ch))
	}
}
