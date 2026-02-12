package slices

import (
	"reflect"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Nil collection",
			args: args{
				source: nil,
			},
			wantResult: []int{},
		},
		{
			name: "Empty collection",
			args: args{
				source: []int{},
			},
			wantResult: []int{},
		},
		{
			name: "Collection with some values",
			args: args{
				source: []int{1, 0, 2, 9, 3, 8, 4, 7, 5, 6},
			},
			wantResult: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Order(tt.args.source); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Order() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestOrderComparable(t *testing.T) {
	type args struct {
		source []time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantResult []time.Time
	}{
		{
			name: "Nil collection",
			args: args{
				source: nil,
			},
			wantResult: []time.Time{},
		},
		{
			name: "Empty collection",
			args: args{
				source: []time.Time{},
			},
			wantResult: []time.Time{},
		},
		{
			name: "Collection with max value",
			args: args{
				source: []time.Time{
					time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 8, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 6, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 4, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			wantResult: []time.Time{
				time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 1, 9, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := OrderComparable(tt.args.source); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("OrderComparable() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestOrderBy(t *testing.T) {
	type args struct {
		source        []string
		valueSelector valueSelector[string, int]
		thenBy        []thenBy[string]
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "Normal test 1",
			args: args{
				source: []string{"grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry"},
				valueSelector: func(s string) int {
					return len(s)
				},
				thenBy: []thenBy[string]{
					ThenBy(func(item string) string { return item }),
				},
			},
			wantResult: []string{"apple", "grape", "mango", "banana", "orange", "blueberry", "raspberry", "passionfruit"},
		},
		{
			name: "Normal test 2",
			args: args{
				source: []string{"grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry"},
				valueSelector: func(s string) int {
					return len(s)
				},
				thenBy: []thenBy[string]{
					ThenDescendingBy(func(item string) string { return item }),
				},
			},
			wantResult: []string{"mango", "grape", "apple", "orange", "banana", "raspberry", "blueberry", "passionfruit"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := OrderBy(tt.args.source, tt.args.valueSelector, tt.args.thenBy...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("OrderBy() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
