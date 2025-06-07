package slices

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/thereisnoplanb/delegate"
)

func TestAggreagate(t *testing.T) {
	type args struct {
		source     []int
		seed       string
		accumulate delegate.Accumulator[int, string]
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "Aggregate test",
			args: args{
				source: []int{1, 2, 3},
				seed:   "Seed ",
				accumulate: func(seed string, value int) (result string) {
					return seed + fmt.Sprintf("%d ", value)
				},
			},
			wantResult: "Seed 1 2 3 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Aggreagate(tt.args.source, tt.args.seed, tt.args.accumulate); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Aggreagate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
