package philosopher

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	//Ch是大管道，里面有5个小管道
	Ch  []chan int
	Sig *sync.WaitGroup
}

func Left(i int) {
	fmt.Println(i, "拿起左叉子")
}
func Right(i int) {
	fmt.Println(i, "拿起右叉子")
}
func Eat(i int) {
	fmt.Println(i, "吃着面")
}

func PutLeft(i int) {
	fmt.Println(i, "放下左叉子")
}

func PutRight(i int) {
	fmt.Println(i, "放下右叉子")
}
func ABS(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func (ph *Philosopher) WantToEat(i int) {
	//采用奇数先拿右边再拿左边，偶数先拿左边再拿右边
	defer func() {
		ph.Ch[i] <- 1
		PutLeft(i)
		ph.Ch[ABS(i-1)%5] <- 1
		PutRight(i)
		ph.Sig.Done()
	}()
	if i%2 == 0 {
		//左
		<-ph.Ch[i]
		Left(i)
		// 右
		<-ph.Ch[ABS(i-1)%5]
		Right(i)
		Eat(i)
	} else {
		// 右
		<-ph.Ch[ABS(i-1)%5]
		Right(i)
		//左
		<-ph.Ch[i]
		Left(i)
		Eat(i)
	}

}

func NewPh() Philosopher {
	return Philosopher{
		Ch:  make([]chan int, 5),
		Sig: &sync.WaitGroup{},
	}
}
