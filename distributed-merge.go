/*
    distributed merge portion of merge sort
    values are received by some other goroutine and appear on channels, this goroutine gets a number of channels as its input and sorts the values coming in; those values are sent to the output channel
*/
package main

import "fmt"
import "sort"

type input struct {
    value int
    source int
}

/*
    function merge takes a slice channels of unsigned integers.  values come in on those channels and they are merge based on value.  note XXX assumptions inline
*/
func merge(sourceChannels []chan uint, output chan uint) {
    // get the number of channels and create a local variable to store the interim values
    var count int = len(sourceChannels)
    var storedInputs = make([]input, count)

    // iterate across the source channels and store the values
    // XXX assumption: channel will send at least one value
    //  if invalid, can make a helper function using code in for loop
    for i := range sourceChannels {
        storedInputs[i] = value{val: <-sourceChannels[i], i}
    }

    // iterate, picking the lowest value in the local slice, and putting it onto the output queue
    // then grab a new value from that input and do it again
    for len(storedInputs) > 0 {
        sort.Slice(storedInputs, func(i, j int) bool { return storedInputs[i].value < storedInputs[j].value }
        output <- storedInputs[0].value
        storedInputs[0].value, ok = <-sourceChannels[storedInputs[0].source]
        if !ok { // the sender has sent a close message; will drop input
            _, storedInputs = storedInputs[0], storedInputs[1:]
        }
    }
}
