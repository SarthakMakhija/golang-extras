# golang-extras

Contains a set of small functions under concurrent/extra package - 
+ Repeat
  + launches a goroutine to repeat a value infinitely and writes these values to an output channel
+ Map
  + launches a goroutine to map values from an incoming channel and writes these mapped values to an output channel
+ Filter
  + launches a goroutine to filter values from an incoming channel and writes these filtered values to an output channel
+ RunningReduce
  + launches a goroutine to perform reduction on the values from an incoming channel and writes these reduced values to an output channel
+ Skip
  + launches a goroutine to skip values from an incoming channel and writes all values other than skipped values to an output channel
+ Reverse
  + launches a goroutine to reverse values from an incoming channel and writes these values to an output channel
+ Take
  + launches a goroutine to take N values from an incoming channel and writes these values to an output channel
+ TakeWhile
  + launches a goroutine to take values from an incoming channel, till the condition is true and writes these values to an output channel
+ DropAll
  + launches a goroutine to drop all values matching an input, from an incoming channel and writes these values to an output channel
+ Merge
  + launches a goroutine to merge incoming channels and writes the merged values to an output channel
+ Tee
  + launches a goroutine to read from an input channel and writes to 2 output channels
