package extra_test

import (
	"github.com/SarthakMakhija/golang-extras/concurrent/extra"
	"math"
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	outputChannel := extra.Repeat(done, func() interface{} {
		return 1
	})

	var elements []interface{}
	for count := 1; count <= 3; count++ {
		elements = append(elements, <-outputChannel)
	}

	expected := []interface{}{1, 1, 1}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Repeat, received %v", expected, elements)
	}
}

func TestMap(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
	}()

	outputChannel := extra.Map(done, inputChannel, func(value interface{}) interface{} {
		return (value.(int)) * 2
	})

	var elements []interface{}
	for mapped := range outputChannel {
		elements = append(elements, mapped)
	}

	expected := []interface{}{2, 4, 6}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Map, received %v", expected, elements)
	}
}

func TestFilter(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
	}()

	outputChannel := extra.Filter(done, inputChannel, func(value interface{}) bool {
		return math.Mod(float64(value.(int)), 2) == 0
	})

	var elements []interface{}
	for mapped := range outputChannel {
		elements = append(elements, mapped)
	}

	expected := []interface{}{2, 4}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Filter, received %v", expected, elements)
	}
}

func TestRunningReduce(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
		inputChannel <- 5
	}()

	outputChannel := extra.RunningReduce(done, inputChannel, 0, func(aggregate interface{}, value interface{}) interface{} {
		return aggregate.(int) + value.(int)
	})

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{1, 3, 6, 10, 15}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from RunningReduce, received %v", expected, elements)
	}
}

func TestSkip(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
	}()

	outputChannel := extra.Skip(done, inputChannel, func(value interface{}) bool {
		return math.Mod(float64(value.(int)), 2) == 0
	})

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{1, 3}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Skip, received %v", expected, elements)
	}
}

func TestReverse(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
	}()

	outputChannel := extra.Reverse(done, inputChannel)

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{4, 3, 2, 1}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Reverse, received %v", expected, elements)
	}
}

func TestTake(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
		inputChannel <- 5
	}()

	outputChannel := extra.Take(done, inputChannel, 4)

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{1, 2, 3, 4}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Take, received %v", expected, elements)
	}
}

func TestTakeWhile(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
		inputChannel <- 5
		inputChannel <- 6
		inputChannel <- 7
	}()

	outputChannel := extra.TakeWhile(done, inputChannel, func(value interface{}) bool {
		return (value.(int)) <= 5
	})

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from TakeWhile, received %v", expected, elements)
	}
}

func TestMerge(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
	}()

	anotherIncomingChannel := make(chan interface{})
	go func() {
		defer close(anotherIncomingChannel)
		anotherIncomingChannel <- 4
		anotherIncomingChannel <- 5
		anotherIncomingChannel <- 6
	}()

	outputChannel := extra.Merge(done, inputChannel, anotherIncomingChannel)

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from Merge, received %v", expected, elements)
	}
}

func TestDropAll(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
		inputChannel <- 4
	}()

	outputChannel := extra.DropAll(done, inputChannel, 4)

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from DropAll, received %v", expected, elements)
	}
}

func TestPipelineUsingRepeatMapTake(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	outputChannel := extra.Take(done,
		extra.Map(done,
			extra.Repeat(done,
				func() interface{} {
					return 2
				},
			),
			func(value interface{}) interface{} {
				return (value.(int)) * 2
			},
		), 4)

	var elements []interface{}
	for element := range outputChannel {
		elements = append(elements, element)
	}

	expected := []interface{}{4, 4, 4, 4}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v from PipelineUsingRepeatMapTake, received %v", expected, elements)
	}
}

func TestTee(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inputChannel := make(chan interface{})
	go func() {
		defer close(inputChannel)
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
	}()

	outputChannel1, outputChannel2 := extra.Tee(done, inputChannel)

	var elementsFromOutputChannel1 []interface{}
	var elementsFromOutputChannel2 []interface{}

	for count := 1; count <= 4; count++ {
		elementsFromOutputChannel1 = append(elementsFromOutputChannel1, <-outputChannel1)
		elementsFromOutputChannel2 = append(elementsFromOutputChannel2, <-outputChannel2)
	}

	expected := []interface{}{1, 2, 3, 4}
	if !reflect.DeepEqual(elementsFromOutputChannel1, expected) {
		t.Fatalf("Expected %v from Tee in OutputChannel1, received %v", expected, elementsFromOutputChannel1)
	}

	if !reflect.DeepEqual(elementsFromOutputChannel1, expected) {
		t.Fatalf("Expected %v from Tee in OutputChannel2, received %v", expected, elementsFromOutputChannel2)
	}
}
