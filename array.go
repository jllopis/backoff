// Array based backoff policy. From Peter Teichman http://blog.gopheracademy.com/advent-2014/backoff/
package backoff

import (
	"fmt"
	"math/rand"
	"time"
)

// ArrayBackoff implements a backoff policy, randomizing its delays
// and saturating at the final value in Millis.
type ArrayBackoff struct {
	Millis       []int
	Tries        int
	LastDuration time.Duration
}

// Default is a backoff policy ranging up to 5 seconds.
var Default = &ArrayBackoff{
	[]int{0, 10, 10, 100, 100, 500, 500, 3000, 3000, 5000},
	0,
	0,
}

func NewArrayBackoff(arr []int) *ArrayBackoff {
	if arr != nil {
		Default = &ArrayBackoff{arr, 0, 0}
	}
	return Default
}

// Duration returns the time duration of the n'th wait cycle in a
// backoff policy. This is b.Millis[n], randomized to avoid thundering
// herds.
func (b *ArrayBackoff) Duration(n int) time.Duration {
	b.Tries = b.Tries + 1
	if n >= len(b.Millis) {
		n = len(b.Millis) - 1
	}

	b.LastDuration = time.Duration(jitter(b.Millis[n])) * time.Millisecond
	return b.LastDuration
}

func (b *ArrayBackoff) String() string {
	return fmt.Sprintf("Try: %d, Duration: %v", b.Tries, b.LastDuration.String())
}

// jitter returns a random integer uniformly distributed in the range
// [0.5 * millis .. 1.5 * millis]
func jitter(millis int) int {
	if millis == 0 {
		return 0
	}

	return millis/2 + rand.Intn(millis)
}
