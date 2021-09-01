package main

type date struct {
	year  int
	month int
	day   int
}

func (d date) LessThan(other date) bool {
	if d.year > other.year {
		return false
	}
	if d.month > other.month {
		return false
	}
	if d.day > other.day {
		return false
	}
	return true
}

// Diff returns the difference in days between d and other
func (d date) Diff(other date) int {
	if d.LessThan(other) {
		d, other = other, d
	}
	diff := 0
	for d.year > other.year {
		diff += d.DaysInYear()
		d.year--
	}
	if d.month > other.month {
		for d.month > other.month {
			diff += d.DaysInMonth()
			d.month--
		}
	} else {
		for other.month > d.month {
			diff -= other.DaysInMonth()
			other.month--
		}
	}
	diff += d.day - other.day
	// You need to calculate the distance in whole days between two dates, counting only the days in between those dates, i.e. 01/01/2001 to 03/01/2001 yields "1".
	return diff - 1
}

func (d date) IsLeapYear() bool {
	return d.year%400 != 0 && d.year%4 == 0
}

func (d date) DaysInMonth() int {
	switch d.month {
	case 2: //February
		if d.IsLeapYear() {
			return 29
		}
		return 28
	case 4, 6, 9, 11: // April, June, September, November
		return 30
	default: // The rest
		return 31
	}
}

func (d date) DaysInYear() int {
	if d.IsLeapYear() {
		return 366
	}
	return 365
}
