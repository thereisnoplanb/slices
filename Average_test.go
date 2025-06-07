package slices

import "testing"

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
			wantResult: 0.0,
			wantErr:    true,
		},
		{
			name: "Collection is nil",
			args: args{
				source: nil,
			},
			wantResult: 0.0,
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
