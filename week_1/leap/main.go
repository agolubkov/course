// Determine if given year is a leap year
package leap

// IsLeapYear returns true if its argument is a leap year
func IsLeapYear(y int) bool {
	switch {
	case y%400 == 0:
		return true
	case y%100 == 0:
		return false
	case y%4 == 0:
		return true
	default:
		return false
	}
}
