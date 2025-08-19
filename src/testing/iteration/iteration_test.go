package iteration

import "testing"

func BenchmarkRepeatRefactored(b *testing.B) {
	for b.Loop() {
		RepeatRefactored("a")
	}
}

func BenchmarkRepeatDumb(b *testing.B) {
	for b.Loop() {
		RepeatDumb("a")
	}
}
