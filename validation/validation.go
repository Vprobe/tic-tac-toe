package validation

func IsEmptyString(s string) bool {
	return len(s) == 0
}

func IsIntOutOfRange(val, min, max int) bool {
	return val < min || val > max
}
