package datastructures

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
	t.Parallel()
	type args struct {
		array []int
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{

		{
			args: args{
				array: []int{1, 2, 3, 4, 5},
				value: 6,
			},
			want:    0,
			wantErr: true,
		},
		{
			args: args{
				array: []int{1, 2, 3, 4},
				value: 3,
			},
			want:    2,
			wantErr: false,
		},
		{
			args: args{
				array: []int{1, 2, 3, 4, 5},
				value: 3,
			},
			want:    2,
			wantErr: false,
		},
		{
			args: args{
				array: []int{1, 2, 3, 4, 5, 6},
				value: 3,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := binarySearch(tt.args.array, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("binarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
