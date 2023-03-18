package slices

//improve https://pkg.go.dev/golang.org/x/exp@v0.0.0-20230224173230-c95f2b4c22f2/slices

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

// ConvertSlice is Convert  T1 -> T2 by convF
func ConvertSlice[T1, T2 any](list []T1, convF func(T1) T2) []T2 {
	var converted []T2
	for _, v := range list {
		converted = append(converted, convF(v))
	}
	return converted
}

// ForeachFind find first occurrence found by f
func ForeachFind[V any, T any](list []V, f func(V) (bool, T)) (found bool, vv T) {
	for _, v := range list {
		found, vv = f(v)
		if found {
			return
		}
	}
	return
}
