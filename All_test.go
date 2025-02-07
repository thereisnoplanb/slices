package slices

import (
	"testing"

	"github.com/thereisnoplanb/generic"
)

func TestAll(t *testing.T) {
	type args struct {
		source    []int
		predicate generic.Predicate[int]
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "All numbers satisfy predicate",
			args: args{
				source: []int{1, 2, 3},
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: true,
		},
		{
			name: "Not all numbers satisfy predicate",
			args: args{
				source: []int{1, 2, 6},
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: false,
		},
		{
			name: "Non number satisfies predicate",
			args: args{
				source: []int{6, 7, 8},
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: false,
		},
		{
			name: "Collection is empty",
			args: args{
				source: []int{},
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: true,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := All(tt.args.source, tt.args.predicate); gotResult != tt.wantResult {
				t.Errorf("All() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
