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
		t.Fatalf("Expected %v from repeat, received %v", expected, elements)
	}
}
