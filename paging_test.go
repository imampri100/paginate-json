package paginatejson

import "testing"

func Test_genTotalPage(t *testing.T) {
	type args struct {
		totalData int
		size      int
	}
	tests := []struct {
		name          string
		args          args
		wantTotalPage int
	}{
		{
			name: "round up on floor condition",
			args: args{
				totalData: 317,
				size:      100,
			},
			wantTotalPage: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotalPage := genTotalPage(tt.args.totalData, tt.args.size); gotTotalPage != tt.wantTotalPage {
				t.Errorf("genTotalPage() = %v, want %v", gotTotalPage, tt.wantTotalPage)
			}
		})
	}
}
