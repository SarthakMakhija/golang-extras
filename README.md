# golang-extras

Contains a set of small functions under concurrent/extra package - 
+ Repeat
  + launches a goroutine to repeat a value infinitely and writes these values to an output channel
+ Map
  + launches a goroutine to map value from an incoming channel and writes these mapped values to an output channel
+ Filter
  + launches a goroutine to filter values from an incoming channel and writes these mapped values to an output channel
+ Take
  + launches a goroutine to take N values from an incoming channel and writes these values to an output channel
+ DropAll
  + launches a goroutine to drop all values, matching an input value from an incoming channel and writes these values to an output channel
+ Merge
  + launches a goroutine to merge incoming channels and writes the merged values to an output channel
