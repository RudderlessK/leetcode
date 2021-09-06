package medium

import "testing"

func Test_centerExpandSolution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := centerExpandSolution(tt.args.s); got != tt.want {
				t.Errorf("centerExpandSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
