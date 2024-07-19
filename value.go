package returndefault

func Value[T comparable](val T, defaultVal T) T {
	var empty T
	if val == empty {
		return defaultVal
	}
	return val
}
