package slices

import (
	"reflect"
	"testing"
)

func TestExcept(t *testing.T) {
	type args struct {
		source []int
		except []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Intersect",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				except: []int{3, 4, 5, 6, 7},
			},
			wantResult: []int{1, 2},
		},
		{
			name: "Intersect nil 1",
			args: args{
				source: nil,
				except: []int{3, 4, 5, 6, 7},
			},
			wantResult: []int{},
		},
		{
			name: "Intersect nil 2",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				except: nil,
			},
			wantResult: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Intersect - no",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				except: []int{6, 7, 8, 9, 0},
			},
			wantResult: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Except(tt.args.source, tt.args.except); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Except() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
