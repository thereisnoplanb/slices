package slices

import (
	"testing"
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
