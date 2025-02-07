package slices

import (
	"testing"

	"github.com/thereisnoplanb/generic"
)

func TestAverage(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
		wantErr    bool
	}{
		{
			name: "Average of some values",
			args: args{
				source: []int{1, 2, 3},
			},
			wantResult: 2.0,
			wantErr:    false,
		},
		{
			name: "Collection is empty",
			args: args{
				source: []int{},
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
			},
			wantResult: 0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Average(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Average() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Average() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestAverageBy(t *testing.T) {
	type SomeType struct {
		Value int
		Other string
	}
	type args struct {
		source        []SomeType
		valueSelector generic.ValueSelector[SomeType, float64]
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
		wantErr    bool
	}{
		{
			name: "Average of some values",
			args: args{
				source: []SomeType{
					{Value: 1, Other: "One"},
					{Value: 2, Other: "Two"},
					{Value: 3, Other: "Three"},
				},
				valueSelector: func(value SomeType) (result float64) {
					return float64(value.Value)
				},
			},
			wantResult: 2.0,
			wantErr:    false,
		},
		{
			name: "Collection is empty",
			args: args{
				source: []SomeType{},
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
			},
			wantResult: 0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := AverageBy(tt.args.source, tt.args.valueSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("AverageBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("AverageBy() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
