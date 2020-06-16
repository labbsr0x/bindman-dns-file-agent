package agent

import (
	"math"
	"time"
)

// backoffWait sleeps thread exponentially longer depending on the trial index
func backoffWait(max uint, triesLeft uint, baseDuration time.Duration) {
	waitSeconds := time.Duration(math.Exp2(float64(max-triesLeft))+1) * baseDuration
	time.Sleep(waitSeconds)
}
