package leap

func IsLeapYear(year int) bool {
	// If it's divisible by 400, it's also divisible by 4 and 100 so this goes 
	// first
	if year % 400 == 0 {
		return true
	} else if year % 100 == 0 {
		// Divisible by 100? Then not leap year
		return false
	} else if year % 4 == 0 {
		return true
	}

	// required return statement at end of function. False by default
	return false
}
