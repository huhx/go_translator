package util

import "testing"

func Test_isChinese(t *testing.T) {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{"include chinese", "I love 中文", true},
		{"all chinese", "中文", true},
		{"all english", "I love you", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinese(tt.text); got != tt.want {
				t.Errorf("isChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}
