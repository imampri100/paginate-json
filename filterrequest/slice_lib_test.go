package filterrequest

import "testing"

func Test_isSlice(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantIsSlice bool
	}{
		{
			name: "want: is slice",
			args: args{
				input: []string{},
			},
			wantIsSlice: true,
		},
		{
			name: "want: is not slice",
			args: args{
				input: "abc",
			},
			wantIsSlice: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsSlice := IsSlice(tt.args.input); gotIsSlice != tt.wantIsSlice {
				t.Errorf("isSlice() = %v, want %v", gotIsSlice, tt.wantIsSlice)
			}
		})
	}
}
