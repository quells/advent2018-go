package strings

func CharsDifferCount(a, b string) int {
	differ := 0
	if len(a) > len(b) {
		return CharsDifferCount(b, a)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differ++
		}
	}
	differ += len(b) - len(a)
	return differ
}
