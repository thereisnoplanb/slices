package slices

import (
	"reflect"
	"testing"
)

func TestCast(t *testing.T) {
	type args struct {
		source []any
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
		wantErr    bool
	}{
		{
			name: "Casting OK",
			args: args{
				source: []any{1, 2, 3},
			},
			wantResult: []int{1, 2, 3},
			wantErr:    false,
		},
		{
			name: "Casting error",
			args: args{
				source: []any{1.0, 2.0, 3.0},
			},
			wantResult: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Cast[int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Cast() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
