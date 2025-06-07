package slices

import (
	"testing"
	"time"

	"github.com/thereisnoplanb/compare"
)

func TestContains(t *testing.T) {
	type args struct {
		source []int
		value  int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Contains value",
			args: args{
				source: []int{1, 2, 3},
				value:  2,
			},
			wantResult: true,
		},
		{
			name: "Do not contain value",
			args: args{
				source: []int{1, 2, 3},
				value:  5,
			},
			wantResult: false,
		},
		{
			name: "Collection is empty",
			args: args{
				source: []int{},
				value:  2,
			},
			wantResult: false,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
				value:  2,
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Contains(tt.args.source, tt.args.value); gotResult != tt.wantResult {
				t.Errorf("Contains() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestContainsFunc(t *testing.T) {
	type args struct {
		source   []int
		value    int
		equality compare.Equality[int]
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Contains value",
			args: args{
				source:   []int{1, 2, 3},
				value:    2,
				equality: compare.Equal[int],
			},
			wantResult: true,
		},
		{
			name: "Do not contain value",
			args: args{
				source:   []int{1, 2, 3},
				value:    5,
				equality: compare.Equal[int],
			},
			wantResult: false,
		},
		{
			name: "Collection is empty",
			args: args{
				source:   []int{},
				value:    2,
				equality: compare.Equal[int],
			},
			wantResult: false,
		},
		{
			name: "Collection is nil",
			args: args{
				source:   nil,
				value:    2,
				equality: compare.Equal[int],
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ContainsFunc(tt.args.source, tt.args.value, tt.args.equality); gotResult != tt.wantResult {
				t.Errorf("ContainsFunc() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestContainsEquatable(t *testing.T) {
	type args struct {
		source []time.Time
		value  time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Contains value",
			args: args{
				source: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				value: time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantResult: true,
		},
		{
			name: "Do not contain value",
			args: args{
				source: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				value: time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantResult: false,
		},
		{
			name: "Collection is empty",
			args: args{
				source: []time.Time{},
				value:  time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantResult: false,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
				value:  time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ContainsEquatable(tt.args.source, tt.args.value); gotResult != tt.wantResult {
				t.Errorf("ContainsEquatable() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
