package gocas

import "sync/atomic"

// GoCas is object that realizes lock-free resource sharing
type GoCas struct {
	flag int32
}

// Wait waits until the shared counter is available. Update the sharing counter if available
func (g *GoCas) Wait() {
	for {
		if g.flag == 0 && atomic.CompareAndSwapInt32(&g.flag, 0, 1) {
			break
		}
	}
}

// Signal signals termination of use of shared counter
func (g *GoCas) Signal() {
	g.flag = 0
}
