package libs

// Index todo
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include todo
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}
