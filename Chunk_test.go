package slices

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		source []int
		size   int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
		wantErr    bool
	}{
		{
			name: "Chunk without remaining part",
			args: args{
				source: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				size:   3,
			},
			wantResult: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			wantErr:    false,
		},
		{
			name: "Chunk with remaining part",
			args: args{
				source: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				size:   3,
			},
			wantResult: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {0}},
			wantErr:    false,
		},
		{
			name: "Chunk empty slice",
			args: args{
				source: []int{},
				size:   3,
			},
			wantResult: [][]int{},
			wantErr:    false,
		},
		{
			name: "Chunk nil slice",
			args: args{
				source: nil,
				size:   3,
			},
			wantResult: [][]int{},
			wantErr:    false,
		},
		{
			name: "Chunk with size below 1",
			args: args{
				source: []int{},
				size:   0,
			},
			wantResult: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Chunk(tt.args.source, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Chunk() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
