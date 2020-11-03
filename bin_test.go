package bin

import (
	"testing"
)

func TestBin(t *testing.T) {
	v1 := int64(64)
	v2 := int32(32)
	v3 := int16(16)
	v4 := float32(32.00)
	v5 := float64(64.00)

	data := Encode(v1, v2, v3, v4, v5)

	if len(data) != 26 {
		t.Fatalf("wrong bytes number")
	}

	v1 = int64(0)
	v2 = int32(0)
	v3 = int16(0)
	v4 = float32(0)
	v5 = float64(0)

	if Decode(data, &v1, &v2, &v3, &v4, &v5) != nil {
		t.Fatal("decode error")
	}

	if v1 != int64(64) || v2 != int32(32) || v3 != int16(16) || v4 != float32(32.00) || v5 != float64(64.00) {
		t.Fatal("decode errors")
	}
}
