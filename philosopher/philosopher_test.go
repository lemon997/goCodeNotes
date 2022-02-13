package philosopher

import (
	"testing"
)

func TestPh(t *testing.T) {
	ph := NewPh()
	for i := 0; i < len(ph.Ch); i++ {
		ph.Ch[i] = make(chan int, 1)
		ph.Ch[i] <- 1
	}
	n := 1
	for i := 0; i < n; i++ {
		ph.Sig.Add(5)
		go ph.WantToEat(0)
		go ph.WantToEat(3)
		go ph.WantToEat(4)
		go ph.WantToEat(2)
		go ph.WantToEat(1)
	}
	ph.Sig.Wait()
}
