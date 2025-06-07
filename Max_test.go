package slices

import (
	"reflect"
	"testing"
	"time"

	"github.com/thereisnoplanb/compare"
)

func TestMax1(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    bool
	}{
		{
			name: "Nil collection",
			args: args{
				source: nil,
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "Empty collection",
			args: args{
				source: []int{},
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "Collection with max value",
			args: args{
				source: []int{1, 5, 7, 2, 4, 3, 6, 9, 8},
			},
			wantResult: 9,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Max(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Max() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMax2(t *testing.T) {
	type args struct {
		source []time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantResult time.Time
		wantErr    bool
	}{
		{
			name: "Nil collection",
			args: args{
				source: nil,
			},
			wantResult: time.Time{},
			wantErr:    true,
		},
		{
			name: "Empty collection",
			args: args{
				source: []time.Time{},
			},
			wantResult: time.Time{},
			wantErr:    true,
		},
		{
			name: "Collection with max value",
			args: args{
				source: []time.Time{
					time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantResult: time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MaxComparable(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxComparable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MaxComparable() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMax3(t *testing.T) {
	type SomeType struct {
		Value int
		Other string
	}
	compareF := func(this SomeType, other SomeType) int {
		return this.Value - other.Value
	}
	type args struct {
		source  []SomeType
		compare compare.Comparison[SomeType]
	}
	tests := []struct {
		name       string
		args       args
		wantResult SomeType
		wantErr    bool
	}{
		{
			name: "Nil collection",
			args: args{
				source:  nil,
				compare: compareF,
			},
			wantResult: SomeType{},
			wantErr:    true,
		},
		{
			name: "Empty collection",
			args: args{
				source:  []SomeType{},
				compare: compareF,
			},
			wantResult: SomeType{},
			wantErr:    true,
		},
		{
			name: "Collection with max value",
			args: args{
				source: []SomeType{
					{Value: 1, Other: "one"},
					{Value: 3, Other: "three"},
					{Value: 2, Other: "two"},
				},
				compare: compareF,
			},
			wantResult: SomeType{Value: 3, Other: "three"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MaxFunc(tt.args.source, tt.args.compare)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MaxFunc() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
