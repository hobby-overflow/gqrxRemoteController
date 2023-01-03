package main

import (
	"testing"
)

func TestSetFreq(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "setFreq 1",
			got:  "1201",
			want: "120100000",
		},
		{
			name: "setFreq 2",
			got:  "120",
			want: "120000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setFreq(tt.got); got != tt.want {
				t.Errorf("setFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}
