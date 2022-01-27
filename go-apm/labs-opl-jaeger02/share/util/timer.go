package util

import (
	"math/rand"
	"time"
)

func DoSomething() {
	time.Sleep(time.Duration(50+rand.Intn(1000)) * time.Millisecond)
}
