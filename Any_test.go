package slices

import "testing"

func TestAny(t *testing.T) {
	type args struct {
		source    []int
		predicate predicate[int]
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Any number satisfies predicate",
			args: args{
				source: []int{1, 2, 3},
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: true,
		},
		{
			name: "Some numbers satisfy predicate",
			args: args{
				source: []int{1, 2, 6},
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: true,
		},
		{
			name: "None number satisfies predicate",
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
			wantResult: false,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
				predicate: func(value int) bool {
					return value < 5
				},
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Any(tt.args.source, tt.args.predicate); gotResult != tt.wantResult {
				t.Errorf("Any() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
