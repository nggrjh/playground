package datastructures

import "testing"

func Test_lint(t *testing.T) {
	type args struct {
		syntax string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				syntax: "( var x = { y: [1, 2, 3] } )",
			},
			wantErr: false,
		},
		{
			args: args{
				syntax: "( var x = { y: [1, 2, 3] }",
			},
			wantErr: true,
		},
		{
			args: args{
				syntax: "( var x = { y: [1, 2, 3] ) }",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lint(tt.args.syntax); (err != nil) != tt.wantErr {
				t.Errorf("lint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
