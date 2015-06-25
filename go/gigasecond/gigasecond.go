package gigasecond

import (
	"time"
)

const TestVersion = 2
var gigaSeconds = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigaSeconds) * time.Second)
}
