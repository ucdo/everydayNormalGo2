package guru

import "testing"

func TestReflectType(t *testing.T) {
	t.Helper()
	reflectType(100)
	reflectType(false)
	reflectType("false")
	reflectType([3]int{1})
}
