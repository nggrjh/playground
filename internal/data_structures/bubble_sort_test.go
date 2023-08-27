package datastructures

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
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
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bubbleSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
