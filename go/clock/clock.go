package clock

import "fmt"

type Clock int
const TestVersion = 1

func New(hour, min int) Clock {
	c := Clock((hour*60 + min) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) Add(add int) Clock {
	c = (c + Clock(add)) % 1440
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
