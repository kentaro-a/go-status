package go-status

import (
	"testing"
)

var masks map[string]BitMask = map[string]BitMask{
	"A": 0b00000001,
	"B": 0b00000010,
	"C": 0b00000100,
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

func TestBits(t *testing.T) {
	st := NewStatus(nil)
	if st.Bits() != "00000000" {
		t.Error()
	}

	st = NewStatus(&Option{masks["A"]})
	if st.Bits() != "00000001" {
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
