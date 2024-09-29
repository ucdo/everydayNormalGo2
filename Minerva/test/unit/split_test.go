package unit

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		want []string
		s    string
		sep  string
	}

	test := &[]args{
		{
			want: []string{"a", "b", "c"},
			s:    "a:b:c",
			sep:  ":",
		},
		{
			want: []string{"a", "b", "c"},
			s:    "aabbabc",
			sep:  "ab",
		},
		{
			want: []string{"abc"},
			s:    "abc",
			sep:  "x",
		},
	}

	for k, v := range *test {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			got := split(v.s, v.sep)
			if ok := reflect.DeepEqual(got, v.want); !ok {
				t.Errorf("split() = %#v, want = %#v", got, v.want)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	b.Log("基准测试：是性能测试")
	for i := 0; i < b.N; i++ {
		split("a:b:c", ":")
	}
}
