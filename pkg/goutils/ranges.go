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

func MinMaxFloat(one float64, two float64) (minValue float64, maxValue float64) {
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

func FloatInRangeInclusive(test float64, one float64, two float64) bool {
	minValue, maxValue := MinMaxFloat(one, two)
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

func FloatInRangeExclusive(test float64, one float64, two float64) bool {
	minValue, maxValue := MinMaxFloat(one, two)
	if test > minValue && test < maxValue {
		return true
	}
	return false
}
