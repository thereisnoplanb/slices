package slices

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	type args struct {
		source []int
		count  int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Nil slice, count < 0",
			args: args{
				source: nil,
				count:  -5,
			},
			wantResult: []int{},
		},
		{
			name: "Nil slice, count = 0",
			args: args{
				source: nil,
				count:  0,
			},
			wantResult: []int{},
		},
		{
			name: "Nil slice, count > 0",
			args: args{
				source: nil,
				count:  5,
			},
			wantResult: []int{},
		},
		{
			name: "Empty slice, count < 0",
			args: args{
				source: []int{},
				count:  -5,
			},
			wantResult: []int{},
		},
		{
			name: "Empty slice, count = 0",
			args: args{
				source: []int{},
				count:  0,
			},
			wantResult: []int{},
		},
		{
			name: "Empty slice, count > 0",
			args: args{
				source: []int{},
				count:  5,
			},
			wantResult: []int{},
		},
		{
			name: "Slice length > 0, count < 0",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  -5,
			},
			wantResult: []int{},
		},
		{
			name: "Slice length > 0, count = 0",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  0,
			},
			wantResult: []int{},
		},
		{
			name: "Slice length > 0, slice length greater than count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  3,
			},
			wantResult: []int{1, 2, 3},
		},
		{
			name: "Slice length > 0, slice length equal count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  5,
			},
			wantResult: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Slice length > 0, slice length less than count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  7,
			},
			wantResult: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Take(tt.args.source, tt.args.count); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Take() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTakeLast(t *testing.T) {
	type args struct {
		source []int
		count  int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Nil slice, count < 0",
			args: args{
				source: nil,
				count:  -5,
			},
			wantResult: []int{},
		},
		{
			name: "Nil slice, count = 0",
			args: args{
				source: nil,
				count:  0,
			},
			wantResult: []int{},
		},
		{
			name: "Nil slice, count > 0",
			args: args{
				source: nil,
				count:  5,
			},
			wantResult: []int{},
		},
		{
			name: "Empty slice, count < 0",
			args: args{
				source: []int{},
				count:  -5,
			},
			wantResult: []int{},
		},
		{
			name: "Empty slice, count = 0",
			args: args{
				source: []int{},
				count:  0,
			},
			wantResult: []int{},
		},
		{
			name: "Empty slice, count > 0",
			args: args{
				source: []int{},
				count:  5,
			},
			wantResult: []int{},
		},
		{
			name: "Slice length > 0, count < 0",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  -5,
			},
			wantResult: []int{},
		},
		{
			name: "Slice length > 0, count = 0",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  0,
			},
			wantResult: []int{},
		},
		{
			name: "Slice length > 0, slice length greater than count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  3,
			},
			wantResult: []int{3, 4, 5},
		},
		{
			name: "Slice length > 0, slice length equal count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  5,
			},
			wantResult: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Slice length > 0, slice length less than count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				count:  7,
			},
			wantResult: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := TakeLast(tt.args.source, tt.args.count); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TakeLast() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
