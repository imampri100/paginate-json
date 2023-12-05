package lib

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
)

func TestIsEmptyFloat64Ptr(t *testing.T) {
	var number *float64
	res := IsEmptyFloat64Ptr(number)
	utils.AssertEqual(t, true, res)

	number = Float64ptr(1)
	res = IsEmptyFloat64Ptr(number)
	utils.AssertEqual(t, false, res)

}

func TestIsEmptyFloat64(t *testing.T) {
	var number float64
	res := IsEmptyFloat64(number)
	utils.AssertEqual(t, true, res)

	number = 1
	res = IsEmptyFloat64(number)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyInt64Ptr(t *testing.T) {
	var number *int64
	res := IsEmptyInt64Ptr(number)
	utils.AssertEqual(t, true, res)

	number = Int64ptr(1)
	res = IsEmptyInt64Ptr(number)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyInt64(t *testing.T) {
	var number int64
	res := IsEmptyInt64(number)
	utils.AssertEqual(t, true, res)

	number = 1
	res = IsEmptyInt64(number)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyIntPtr(t *testing.T) {
	var number *int
	res := IsEmptyIntPtr(number)
	utils.AssertEqual(t, true, res)

	number = Intptr(1)
	res = IsEmptyIntPtr(number)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyInt(t *testing.T) {
	var number int
	res := IsEmptyInt(number)
	utils.AssertEqual(t, true, res)

	number = 1
	res = IsEmptyInt(number)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyStrPtr(t *testing.T) {
	var str *string
	res := IsEmptyStrPtr(str)
	utils.AssertEqual(t, true, res)

	str = Strptr("1")
	res = IsEmptyStrPtr(str)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyStr(t *testing.T) {
	var str string
	res := IsEmptyStr(str)
	utils.AssertEqual(t, true, res)

	str = "1"
	res = IsEmptyStr(str)
	utils.AssertEqual(t, false, res)
}

func TestIsFalsyBoolPtr(t *testing.T) {
	var isBool *bool
	res := IsFalsyBoolPtr(isBool)
	utils.AssertEqual(t, true, res)

	isBool = Boolptr(true)
	res = IsFalsyBoolPtr(isBool)
	utils.AssertEqual(t, false, res)
}

func TestIsFalsyBool(t *testing.T) {
	var isBool bool
	res := IsFalsyBool(isBool)
	utils.AssertEqual(t, true, res)

	isBool = true
	res = IsFalsyBool(isBool)
	utils.AssertEqual(t, false, res)

}

func TestIsEmptyUUIDPtr(t *testing.T) {
	var id *uuid.UUID
	res := IsEmptyUUIDPtr(id)
	utils.AssertEqual(t, true, res)

	id = UUIDPtr(uuid.New())
	res = IsEmptyUUIDPtr(id)
	utils.AssertEqual(t, false, res)
}

func TestIsEmptyUUID(t *testing.T) {
	var id uuid.UUID
	res := IsEmptyUUID(id)
	utils.AssertEqual(t, true, res)

	id = uuid.New()
	res = IsEmptyUUID(id)
	utils.AssertEqual(t, false, res)
}

func TestIsZeroTimePtr(t *testing.T) {
	now := time.Now()

	type args struct {
		moment *time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantIsZero bool
	}{
		// Add test cases.
		{
			name: "must zero",
			args: args{
				moment: &time.Time{},
			},
			wantIsZero: true,
		},
		{
			name: "not zero",
			args: args{
				moment: &now,
			},
			wantIsZero: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsZero := IsZeroTimePtr(tt.args.moment); gotIsZero != tt.wantIsZero {
				t.Errorf("IsZeroTimePtr() = %v, want %v", gotIsZero, tt.wantIsZero)
			}
		})
	}
}

func TestIsZeroStrfmtTimePtr(t *testing.T) {
	var date *strfmt.DateTime

	res := IsZeroStrfmtTimePtr(date)
	utils.AssertEqual(t, true, res)

	now := time.Now()
	date = DateTimeptr((strfmt.DateTime)(now))
	res = IsZeroStrfmtTimePtr(date)
	utils.AssertEqual(t, false, res)

}

func TestIsZeroStrfmtTime(t *testing.T) {
	var date strfmt.DateTime

	res := IsZeroStrfmtTime(date)
	utils.AssertEqual(t, true, res)

	now := time.Now()
	date = (strfmt.DateTime)(now)
	res = IsZeroStrfmtTime(date)
	utils.AssertEqual(t, false, res)
}

func TestIsSimilarStringPattern(t *testing.T) {
	isSimilar := IsSimilarStringPattern(`[a-z]+`, "test abc")
	utils.AssertEqual(t, true, isSimilar, "validate response")

	isSimilar = IsSimilarStringPattern(`[0-9]+`, "test abc")
	utils.AssertEqual(t, false, isSimilar, "validate response")
}
