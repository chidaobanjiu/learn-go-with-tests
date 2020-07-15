package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func ExampleRepeat() {
	res := Repeat("b", 3)
	fmt.Println(res)
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func TestRepeatWithStrings(t *testing.T) {
	repeated := RepeatWithStrings("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func ExampleRepeatWithStrings() {
	res := RepeatWithStrings("b", 3)
	fmt.Println(res)
	// Output: bbb
}

func BenchmarkRepeatWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithStrings("a", 10)
	}
}
