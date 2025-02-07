package slices

import (
	"reflect"
	"testing"

	"github.com/thereisnoplanb/generic"
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

func TestSumBy(t *testing.T) {
	type SomeType struct {
		Value int
		Other string
	}
	valueSelector := func(value SomeType) int {
		return value.Value
	}
	type args struct {
		source        []SomeType
		valueSelector generic.ValueSelector[SomeType, int]
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "Nil collection",
			args: args{
				source:        nil,
				valueSelector: valueSelector,
			},
			wantResult: 0,
		},
		{
			name: "Empty collection",
			args: args{
				source:        []SomeType{},
				valueSelector: valueSelector,
			},
			wantResult: 0,
		},
		{
			name: "Collection with some values",
			args: args{
				source: []SomeType{
					{Value: 1, Other: "one"},
					{Value: 2, Other: "two"},
					{Value: 3, Other: "three"},
					{Value: 4, Other: "four"},
					{Value: 5, Other: "five"},
					{Value: 6, Other: "six"},
					{Value: 7, Other: "seven"},
					{Value: 8, Other: "eight"},
					{Value: 9, Other: "nine"},
					{Value: 0, Other: "zero"},
				},
				valueSelector: valueSelector,
			},
			wantResult: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SumBy(tt.args.source, tt.args.valueSelector); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SumBy() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
