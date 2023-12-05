package filterrequest

import (
	"reflect"
	"testing"
)

func TestGenListRawWhereClause(t *testing.T) {
	type args struct {
		iface           []interface{}
		wrapperPosition PositionWhereClause
	}
	tests := []struct {
		name    string
		args    args
		wantLwc *ListRawWhereClause
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				iface:           []interface{}{"current", "=", "proposal"},
				wrapperPosition: PositionWhereClause{},
			},
			wantLwc: &ListRawWhereClause{
				{
					rawData: []interface{}{"current", "=", "proposal"},
					position: PositionWhereClause{
						slice:  []interface{}{"current", "=", "proposal"},
						level:  0,
						index:  0,
						length: 3,
					},
				},
			},
		},
		{
			name: "success",
			args: args{
				iface: []interface{}{
					[]interface{}{"current", "=", "proposal"},
				},
				wrapperPosition: PositionWhereClause{},
			},
			wantLwc: &ListRawWhereClause{
				{
					rawData: []interface{}{"current", "=", "proposal"},
					position: PositionWhereClause{
						slice:  []interface{}{"current", "=", "proposal"},
						level:  1,
						index:  0,
						length: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLwc, err := GenListRawWhereClause(tt.args.iface, tt.args.wrapperPosition)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenListRawWhereClause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLwc, tt.wantLwc) {
				t.Errorf("GenListRawWhereClause() = %v, want %v", gotLwc, tt.wantLwc)
			}
		})
	}
}
