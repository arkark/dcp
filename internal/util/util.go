package util

func Contains(xs []string, y string) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func IfElse(cond bool, trueValue string, falseValue string) string {
	if cond {
		return trueValue
	} else {
		return falseValue
	}
}
