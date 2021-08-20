package extra_test

import (
	"github.com/SarthakMakhija/golang-extras/concurrent/extra"
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

	incomingChannel := make(chan interface{})
	go func() {
		defer close(incomingChannel)
		incomingChannel <- 1
		incomingChannel <- 2
		incomingChannel <- 3
	}()

	outputChannel := extra.Map(done, incomingChannel, func(value interface{}) interface{} {
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
