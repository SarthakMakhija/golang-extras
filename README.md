# golang-extras

Contains a set of small functions under concurrent/extra package - 
+ Repeat
  + launches a goroutine to repeat a value infinitely and writes these values to an output channel
+ Map
  + launches a goroutine to map values from an input channel and writes these mapped values to an output channel
+ Filter
  + launches a goroutine to filter values from an input channel and writes these filtered values to an output channel
+ RunningReduce
  + launches a goroutine to perform reduction on the values from an input channel and writes these reduced values to an output channel
+ Skip
  + launches a goroutine to skip values from an input channel and writes all values other than skipped values to an output channel
+ Reverse
  + launches a goroutine to reverse values from an input channel and writes these values to an output channel
+ Take
  + launches a goroutine to take N values from an input channel and writes these values to an output channel
+ TakeWhile
  + launches a goroutine to take values from an input channel, till the condition is true and writes these values to an output channel
+ DropAll
  + launches a goroutine to drop all values matching an input, from an input channel and writes these values to an output channel
+ Merge
  + launches a goroutine to merge input channels and writes the merged values to an output channel
+ Tee
  + launches a goroutine to read from an input channel and writes to 2 output channels


# Usage

**Repeat**

```golang
  
done := make(chan interface{})
defer close(done)

outputChannel := extra.Repeat(done, func() interface{} {
  return 1
})

var elements []interface{}
for count := 1; count <= 3; count++ {
  elements = append(elements, <-outputChannel)
}
```

**Map**

```golang
  
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
```

**Creating a pipeline using Repeat, Map and Take**

```golang
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
```

