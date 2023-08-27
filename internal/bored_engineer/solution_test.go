package playground_test

import (
	"testing"

	playground "github.com/nggrjh/playground/internal/bored_engineer"
)

func TestBoredEngineerWorkingHours(t *testing.T) {
	type args struct {
		tasks []int
		gap   int
	}
	tests := map[string]struct {
		args args
		want int
	}{
		"case_1": {
			args: args{
				tasks: []int{1, 2, 3, 1, 1, 2, 1},
				gap:   2,
			},
			want: 10,
		},
		"case_2": {
			args: args{
				tasks: []int{1, 2, 3, 4, 5, 6, 7, 8},
				gap:   2,
			},
			want: 8,
		},
		"case_3": {
			args: args{
				tasks: []int{1, 1, 1, 1, 1, 1, 1, 1},
				gap:   2,
			},
			want: 22,
		},
	}
	for name, test := range tests {
		nm := name
		tt := test
		t.Run(nm, func(t *testing.T) {
			if got := playground.BoredEngineerWorkingHours(tt.args.tasks, tt.args.gap); got != tt.want {
				t.Errorf("BoredEngineerWorkingHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoredEngineerRearrangeWorkingHours(t *testing.T) {
	type args struct {
		tasks []int
		gap   int
	}
	tests := map[string]struct {
		args args
		want int
	}{
		// "case_1": {
		// 	args: args{
		// 		tasks: []int{1, 2, 3, 1, 1, 2, 1},
		// 		gap:   2,
		// 	},
		// 	want: 1000000,
		// },
		// "case_3": {
		// 	args: args{
		// 		tasks: []int{1, 1, 1, 1, 1, 1, 1, 1},
		// 		gap:   2,
		// 	},
		// 	want: 22,
		// },
	}
	for name, test := range tests {
		nm := name
		tt := test
		t.Run(nm, func(t *testing.T) {
			if got := playground.BoredEngineerRearrangeWorkingHours(tt.args.tasks, tt.args.gap); got != tt.want {
				t.Errorf("BoredEngineerRearrangeWorkingHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
