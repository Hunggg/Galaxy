package common

func MaxInt64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func MinInt64(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}
