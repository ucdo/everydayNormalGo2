package guru

import "testing"

func TestReflectType(t *testing.T) {
	t.Helper()
	reflectType(100)
	reflectType(false)
	reflectType("false")
	reflectType([3]int{1})
}

func TestModify(t *testing.T) {
	a := 100
	reflectType(a)
	t.Log(a)
}
