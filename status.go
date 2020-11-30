package gostatus

import (
	"fmt"
)

const (
	BIT_0 BitMask = 1 << iota
	BIT_1
	BIT_2
	BIT_3
	BIT_4
	BIT_5
	BIT_6
	BIT_7
)

type BitMask uint8

type Status struct {
	Value uint8
}

type Option struct {
	BitMask BitMask
}

func NewStatus(o *Option) *Status {
	_f := &Status{
		Value: 0,
	}
	if o != nil {
		_f.On(o.BitMask)
	}
	return _f
}

func (f *Status) Bits() string {
	return fmt.Sprintf("%08b", f.Value)
}

func (f *Status) IsOn(b BitMask) bool {
	bit := f.Value & uint8(b)
	return bit > 0
}

func (f *Status) On(b BitMask) {
	f.Value = f.Value | uint8(b)
}

func (f *Status) Off(b BitMask) {
	f.Value = f.Value & ^uint8(b)
}
