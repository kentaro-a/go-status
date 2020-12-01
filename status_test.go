package gostatus

import (
	"testing"
)

var masks map[string]Bits = map[string]Bits{
	"A": BIT_0,
	"B": BIT_1,
	"C": BIT_2,
	"D": BIT_31,
}

func TestNewStatus(t *testing.T) {
	st := NewStatus(nil)
	if st.Value != 0 {
		t.Error()
	}

	st = NewStatus(&Option{masks["A"]})
	if st.Value == 0 {
		t.Error()
	}
}

func TestGetAllOnBits(t *testing.T) {
	var expected Bits = 0b11111111111111111111111111111111
	if expected != GetAllOnBits() {
		t.Error()
	}
}

func TestGetAllOffBits(t *testing.T) {
	var allon Bits = 0b11111111111111111111111111111111
	expected := ^allon
	if expected != GetAllOffBits() {
		t.Error()
	}
}

func TestSetBits(t *testing.T) {
	st := NewStatus(nil)
	bits := Bits(0b10000000000000000000000000000001)
	st.SetBits(bits)
	if st.GetBits() != bits {
		t.Error()
	}
}

func TestSetBitsString(t *testing.T) {
	st := NewStatus(nil)
	bits_str := "10101010101010101010101010101011"
	err := st.SetBitsString(bits_str)
	if err != nil {
		t.Error(err)
	}
	if st.GetBitsString() != bits_str {
		t.Error(st.GetBitsString())
	}

	bits_str_min := "00000000000000000000000000000000"
	err = st.SetBitsString(bits_str_min)
	if err != nil {
		t.Error(err)
	}
	if st.GetBitsString() != bits_str_min {
		t.Error(st.GetBitsString())
	}

	bits_str_max := "11111111111111111111111111111111"
	err = st.SetBitsString(bits_str_max)
	if err != nil {
		t.Error(err)
	}
	if st.GetBitsString() != bits_str_max {
		t.Error(st.GetBitsString())
	}

	bits_str_overflow := "111111111111111111111111111111111"
	err = st.SetBitsString(bits_str_overflow)
	if err == nil {
		t.Error()
	}
}

func TestGetBits(t *testing.T) {
	st := NewStatus(nil)

	if st.GetBits() != Bits(0b00000000000000000000000000000000) {
		t.Error()
	}

	st = NewStatus(&Option{masks["A"]})
	if st.GetBits() != Bits(0b00000000000000000000000000000001) {
		t.Error()
	}

	st = NewStatus(&Option{masks["D"] | masks["B"]})
	if st.GetBits() != Bits(0b10000000000000000000000000000010) {
		t.Error()
	}
}

func TestGetBitsUint32(t *testing.T) {
	st := NewStatus(nil)
	bits := Bits(0b00000000000000000000000000000001)
	st.SetBits(bits)
	if st.GetBitsUint32() != uint32(bits) {
		t.Error()
	}

}

func TestGetBitsString(t *testing.T) {
	st := NewStatus(nil)
	if st.GetBitsString() != "00000000000000000000000000000000" {
		t.Error()
	}

	st = NewStatus(&Option{masks["A"]})
	if st.GetBitsString() != "00000000000000000000000000000001" {
		t.Error()
	}

	st = NewStatus(&Option{masks["D"] | masks["B"]})
	if st.GetBitsString() != "10000000000000000000000000000010" {
		t.Error()
	}
}

func TestIsOn(t *testing.T) {
	st := NewStatus(&Option{masks["A"]})
	if st.IsOn(masks["A"]) != true {
		t.Error()
	}
	if st.IsOn(masks["B"]) != false {
		t.Error()
	}
	if st.IsOn(masks["C"]) != false {
		t.Error()
	}
	st.On(masks["B"])
	// A:on, B:on, C:off

	if st.IsOn(masks["A"]|masks["B"]) != true {
		t.Error()
	}
	if st.IsOn(masks["A"]|masks["C"]) != false {
		t.Error()
	}
	if st.IsOn(masks["B"]|masks["C"]) != false {
		t.Error()
	}
}

func TestIsOff(t *testing.T) {
	st := NewStatus(&Option{masks["A"]})
	if st.IsOff(masks["A"]) != false {
		t.Error()
	}
	if st.IsOff(masks["B"]) != true {
		t.Error()
	}
	if st.IsOff(masks["C"]) != true {
		t.Error()
	}
	st.On(masks["B"])
	// A:on, B:on, C:off

	if st.IsOff(masks["A"]) != false {
		t.Error()
	}
	if st.IsOff(masks["B"]) != false {
		t.Error()
	}
	if st.IsOff(masks["C"]) != true {
		t.Error()
	}
	if st.IsOff(masks["A"]|masks["B"]) != false {
		t.Error()
	}
	if st.IsOff(masks["A"]|masks["C"]) != false {
		t.Error()
	}
	if st.IsOff(masks["B"]|masks["C"]) != false {
		t.Error()
	}
}

func TestOn(t *testing.T) {
	st := NewStatus(nil)
	st.On(masks["B"])
	if st.IsOn(masks["A"]) != false {
		t.Error()
	}
	if st.IsOn(masks["B"]) != true {
		t.Error()
	}
	if st.IsOn(masks["C"]) != false {
		t.Error()
	}

	st.On(masks["A"] | masks["C"])
	if st.IsOn(masks["A"]) != true {
		t.Error()
	}
	if st.IsOn(masks["B"]) != true {
		t.Error()
	}
	if st.IsOn(masks["C"]) != true {
		t.Error()
	}
}

func TestOff(t *testing.T) {
	st := NewStatus(&Option{masks["A"] | masks["B"]})
	st.Off(masks["A"])
	if st.IsOn(masks["A"]) != false {
		t.Error()
	}
	if st.IsOn(masks["B"]) != true {
		t.Error()
	}
	if st.IsOn(masks["C"]) != false {
		t.Error()
	}

	st.Off(masks["A"] | masks["B"] | masks["C"])
	if st.IsOn(masks["A"]) != false {
		t.Error()
	}
	if st.IsOn(masks["B"]) != false {
		t.Error()
	}
	if st.IsOn(masks["C"]) != false {
		t.Error()
	}
}
