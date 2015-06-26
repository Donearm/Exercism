package raindrops

import (
	"fmt"
	"strconv"
)

func Convert(n int) string {
	var ret_string string
	if ((n % 3) == 0) { ret_string = "Pling" }
	if ((n % 5) == 0) { ret_string = fmt.Sprint(ret_string, "Plang") }
	if ((n % 7) == 0) { ret_string = fmt.Sprint(ret_string, "Plong") }

	if len(ret_string) == 0 {
		ret_string = strconv.Itoa(n)
	}

	return ret_string
}

