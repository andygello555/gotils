package concurrency

// In-Out infinite channels that don't block when written to.
//
// This is from https://medium.com/capital-one-tech/building-an-unbounded-channel-in-go-789e175cd2cd.
func InOut() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})

	go func() {
		var inQueue []interface{}

		// Temp function which returns the out channel to write to
		// This is done to avoid writing nils to the out channel
		outCh := func() chan interface{} {
			if len(inQueue) == 0 {
				return nil
			}
			return out
		}

		// Returns the head of the input queue if the queue is not empty otherwise it returns nil
		curVal := func() interface {} {
			if len(inQueue) == 0 {
				return nil
			}
			return inQueue[0]
		}

		for len(inQueue) > 0 || in != nil {
			select {
			// Read from input channel if we can
			case v, ok := <-in:
				if !ok {
					// If input channel is empty then we set input to a nil channel so we don't read anything more
					in = nil
				} else {
					// We append the input to the queue to be written to out
					inQueue = append(inQueue, v)
				}
			// If there is a value in the queue to write to out then write
			case outCh() <- curVal():
				// We pop off the head of the queue
				inQueue = inQueue[1:]
			}
		}
		close(out)
	}()
	return in, out
}
