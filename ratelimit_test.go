package ratelimit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lxstvayne/ratelimit"
)

func Test_Ratelimit(t *testing.T) {
	limiter := ratelimit.New(10, ratelimit.WithSlack(10))

	limiter.Take()
	time.Sleep(time.Millisecond * 500)

	for i := 0; i < 10; i++ {
		limiter.Take()
		fmt.Printf("Taken time: %s\n", time.Now())
		fmt.Printf("Next sleep duration: %s\n", limiter.NextSleepDuration())
	}
}
