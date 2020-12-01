package gostatus

import (
	"errors"
	"fmt"
	"strconv"
)

type Bits uint32

type Status struct {
	Value Bits
}

type Option struct {
	Bits Bits
}

func NewStatus(o *Option) *Status {
	_f := &Status{
		Value: 0,
	}
	if o != nil {
		_f.On(o.Bits)
	}
	return _f
}

func GetAllOnBits() Bits {
	return 0b11111111111111111111111111111111
}

func GetAllOffBits() Bits {
	all_on := GetAllOnBits()
	all_off := ^all_on
	return all_off
}

func (f *Status) SetBits(b Bits) {
	f.Value = b
}

func (f *Status) SetBitsString(s string) error {
	var b Bits
	var max Bits = GetAllOnBits()
	input, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return err
	}
	if input > uint64(max) {
		return errors.New("Input bits string is overflow.")
	}
	b = Bits(input)
	f.Value = b
	return nil
}

func (f *Status) GetBits() Bits {
	return f.Value
}

func (f *Status) GetBitsUint32() uint32 {
	return uint32(f.Value)
}

func (f *Status) GetBitsString() string {
	return fmt.Sprintf("%032b", f.Value)
}

func (f *Status) IsOn(b Bits) bool {
	bits := f.Value & b
	return bits == b
}

func (f *Status) IsOff(b Bits) bool {
	bits := f.Value & b
	return bits == 0
}

func (f *Status) On(b Bits) {
	f.Value = f.Value | b
}

func (f *Status) Off(b Bits) {
	f.Value = f.Value & ^b
}
