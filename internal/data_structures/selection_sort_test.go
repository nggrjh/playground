package datastructures

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	t.Parallel()
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				array: []int{4, 2, 7, 1, 3},
			},
			want: []int{1, 2, 3, 4, 7},
		},
		{
			args: args{
				array: []int{7, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 7},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := selectionSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
