package slices

import (
	"testing"
	"time"

	"github.com/thereisnoplanb/compare"
)

func TestContainsAny(t *testing.T) {
	type args struct {
		source []int
		values []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Source contains any value",
			args: args{
				source: []int{1, 2, 3},
				values: []int{1, 5},
			},
			wantResult: true,
		},
		{
			name: "Source does not contain any value",
			args: args{
				source: []int{1, 2, 3},
				values: []int{7, 5},
			},
			wantResult: false,
		},
		{
			name: "Source is empty",
			args: args{
				source: []int{},
				values: []int{1, 2},
			},
			wantResult: false,
		},
		{
			name: "Source is nil",
			args: args{
				source: nil,
				values: []int{1, 2},
			},
			wantResult: false,
		},
		{
			name: "Values is empty",
			args: args{
				source: []int{1, 2, 3},
				values: []int{},
			},
			wantResult: false,
		},
		{
			name: "Values is nil",
			args: args{
				source: []int{1, 2, 3},
				values: nil,
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ContainsAny(tt.args.source, tt.args.values...); gotResult != tt.wantResult {
				t.Errorf("ContainsAny() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	type args struct {
		source   []int
		equality compare.Equality[int]
		values   []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Source contains any value",
			args: args{
				source: []int{1, 2, 3},
				equality: func(x, y int) bool {
					return x == y
				},
				values: []int{1, 5},
			},
			wantResult: true,
		},
		{
			name: "Source does not contain any value",
			args: args{
				source: []int{1, 2, 3},
				equality: func(x, y int) bool {
					return x == y
				},
				values: []int{7, 5},
			},
			wantResult: false,
		},
		{
			name: "Source is empty",
			args: args{
				source: []int{},
				equality: func(x, y int) bool {
					return x == y
				},
				values: []int{1, 2},
			},
			wantResult: false,
		},
		{
			name: "Source is nil",
			args: args{
				source: nil,
				equality: func(x, y int) bool {
					return x == y
				},
				values: []int{1, 2},
			},
			wantResult: false,
		},
		{
			name: "Values is empty",
			args: args{
				source: []int{1, 2, 3},
				equality: func(x, y int) bool {
					return x == y
				},
				values: []int{},
			},
			wantResult: false,
		},
		{
			name: "Values is nil",
			args: args{
				source: []int{1, 2, 3},
				equality: func(x, y int) bool {
					return x == y
				},
				values: nil,
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ContainsAnyFunc(tt.args.source, tt.args.equality, tt.args.values...); gotResult != tt.wantResult {
				t.Errorf("ContainsAnyFunc() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestContainsAnyEquatable(t *testing.T) {
	type args struct {
		source []time.Time
		values []time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Source contains any value",
			args: args{
				source: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				values: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantResult: true,
		},
		{
			name: "Source does not contain any value",
			args: args{
				source: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				values: []time.Time{
					time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantResult: false,
		},
		{
			name: "Source is empty",
			args: args{
				source: []time.Time{},
				values: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantResult: false,
		},
		{
			name: "Source is nil",
			args: args{
				source: nil,
				values: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantResult: false,
		},
		{
			name: "Values is empty",
			args: args{
				source: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				values: []time.Time{},
			},
			wantResult: false,
		},
		{
			name: "Values is nil",
			args: args{
				source: []time.Time{
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				values: nil,
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ContainsAnyEquatable(tt.args.source, tt.args.values...); gotResult != tt.wantResult {
				t.Errorf("ContainsAnyEquatable() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
