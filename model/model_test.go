package model

import (
	"testing"
)

func TestAddExtraSpaceIfExist(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", " a"},
		{" a", "  a"},
		{" ", "  "},
	}
	for i, tt := range tests {
		if got := addExtraSpaceIfExist(tt.input); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
