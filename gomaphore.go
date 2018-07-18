package gomaphore

import "sync/atomic"

// Gomaphore is semaphore object
type Gomaphore struct {
	flag int32
}

// Wait for a semaphore and set flag on
func (g *Gomaphore) Wait() {
	for {
		if g.flag == 0 && atomic.CompareAndSwapInt32(&g.flag, 0, 1) {
			break
		}
	}
}

// Signal signals a semaphore and set flag off
func (g *Gomaphore) Signal() {
	g.flag = 0
}
