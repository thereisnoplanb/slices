package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "Nil collection",
			args: args{
				source: nil,
			},
			wantResult: 0,
		},
		{
			name: "Empty collection",
			args: args{
				source: []int{},
			},
			wantResult: 0,
		},
		{
			name: "Collection with some values",
			args: args{
				source: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			wantResult: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Sum(tt.args.source); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Sum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
