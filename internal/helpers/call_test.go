package helpers

import "testing"

func TestCalculateBlockCount(t *testing.T) {
	type args struct {
		duration int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "duration divisible for block_size",
			args: args{
				duration: 60,
			},
			want: 2,
		},
		{
			name: "duration not divisible for block_size",
			args: args{
				duration: 70,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateBlockCount(tt.args.duration); got != tt.want {
				t.Errorf("CalculateBlockCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
