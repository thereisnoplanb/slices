package slices

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Intersect",
			args: args{
				first:  []int{1, 2, 3, 4, 5},
				second: []int{3, 4, 5, 6, 7},
			},
			wantResult: []int{3, 4, 5},
		},
		{
			name: "Intersect no distinct elements",
			args: args{
				first:  []int{1, 1, 2, 2, 3},
				second: []int{1, 2, 2, 3, 3},
			},
			wantResult: []int{1, 2, 3},
		},
		{
			name: "Intersect nil 1",
			args: args{
				first:  nil,
				second: []int{3, 4, 5, 6, 7},
			},
			wantResult: []int{},
		},
		{
			name: "Intersect nil 2",
			args: args{
				first:  []int{1, 2, 3, 4, 5},
				second: nil,
			},
			wantResult: []int{},
		},
		{
			name: "Intersect - no",
			args: args{
				first:  []int{1, 2, 3, 4, 5},
				second: []int{6, 7, 8, 9, 0},
			},
			wantResult: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Intersect(tt.args.first, tt.args.second); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Intersect() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
