package slices

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		collection1 []int
		collection2 []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Intersect",
			args: args{
				collection1: []int{1, 2, 3, 4, 5},
				collection2: []int{3, 4, 5, 6, 7},
			},
			wantResult: []int{3, 4, 5},
		},
		{
			name: "Intersect nil 1",
			args: args{
				collection1: nil,
				collection2: []int{3, 4, 5, 6, 7},
			},
			wantResult: []int{},
		},
		{
			name: "Intersect nil 2",
			args: args{
				collection1: []int{1, 2, 3, 4, 5},
				collection2: nil,
			},
			wantResult: []int{},
		},
		{
			name: "Intersect - no",
			args: args{
				collection1: []int{1, 2, 3, 4, 5},
				collection2: []int{6, 7, 8, 9, 0},
			},
			wantResult: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Intersect(tt.args.collection1, tt.args.collection2); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Intersect() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
