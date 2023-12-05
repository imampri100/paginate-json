package filterrequest

import (
	"testing"
)

func TestGenValueKind(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantVkd ValueKind
		wantErr bool
	}{
		{
			name: "input nil, want nil",
			args: args{
				input: nil,
			},
			wantVkd: NilValKind,
			wantErr: false,
		},
		{
			name: "input slice int, want slice int",
			args: args{
				input: []interface{}{90},
			},
			wantVkd: SliceFloat64ValKind,
			wantErr: false,
		},
		{
			name: "input slice float64, want slice float64",
			args: args{
				input: []interface{}{90.10},
			},
			wantVkd: SliceFloat64ValKind,
			wantErr: false,
		},
		{
			name: "input slice string, want slice string",
			args: args{
				input: []interface{}{"90"},
			},
			wantVkd: SliceStringValKind,
			wantErr: false,
		},
		{
			name: "input int, want int",
			args: args{
				input: 90,
			},
			wantVkd: Float64ValKind,
			wantErr: false,
		},
		{
			name: "input float64, want float64",
			args: args{
				input: 90.10,
			},
			wantVkd: Float64ValKind,
			wantErr: false,
		},
		{
			name: "input string, want string",
			args: args{
				input: "90",
			},
			wantVkd: StringValKind,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVkd, err := GenValueKind(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenValueKind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVkd != tt.wantVkd {
				t.Errorf("GenValueKind() = %v, want %v", gotVkd, tt.wantVkd)
			}
		})
	}
}

func TestValueKind_IsMatch(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name        string
		vkd         ValueKind
		args        args
		wantIsMatch bool
	}{
		{
			name: "",
			vkd:  SliceStringValKind,
			args: args{
				input: []interface{}{"abc"},
			},
			wantIsMatch: true,
		},
		{
			name: "",
			vkd:  SliceStringValKind,
			args: args{
				input: []interface{}{10},
			},
			wantIsMatch: false,
		},
		{
			name: "",
			vkd:  NilValKind,
			args: args{
				input: []interface{}{10},
			},
			wantIsMatch: false,
		},
		{
			name: "",
			vkd:  NilValKind,
			args: args{
				input: nil,
			},
			wantIsMatch: true,
		},
		{
			name: "",
			vkd:  SliceStringValKind,
			args: args{
				input: nil,
			},
			wantIsMatch: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsMatch := tt.vkd.IsMatch(tt.args.input); gotIsMatch != tt.wantIsMatch {
				t.Errorf("ValueKind.IsMatch() = %v, want %v", gotIsMatch, tt.wantIsMatch)
			}
		})
	}
}
