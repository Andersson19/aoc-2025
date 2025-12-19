package util

// Source - https://stackoverflow.com/a
// Posted by T. Claverie, modified by community. See post 'Timeline' for change history
// Retrieved 2025-12-17, License - CC BY-SA 4.0
func Remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func AreEqual[K comparable](a, b []K) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
