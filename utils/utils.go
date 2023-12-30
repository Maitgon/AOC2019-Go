package utils

func Map[T, F any](vals []T, f func(T) F) []F {
	res := make([]F, len(vals))
	for i, val := range vals {
		res[i] = f(val)
	}
	return res
}
