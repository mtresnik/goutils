package goutils

func MinMaxInt(one int, two int) (minValue int, maxValue int) {
	if one < two {
		minValue = one
		maxValue = two
	} else {
		minValue = two
		maxValue = one
	}
	return
}

func IntInRangeInclusive(test int, one int, two int) bool {
	minValue, maxValue := MinMaxInt(one, two)
	if test >= minValue && test <= maxValue {
		return true
	}
	return false
}

func IntInRangeExclusive(test int, one int, two int) bool {
	minValue, maxValue := MinMaxInt(one, two)
	if test > minValue && test < maxValue {
		return true
	}
	return false
}
