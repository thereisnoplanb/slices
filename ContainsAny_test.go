package slices

import "testing"

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
