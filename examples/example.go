package main

import (
	"fmt"

	"github.com/kentaro-a/gostatus"
)

func main() {
	masks := map[string]gostatus.Bits{
		"A": gostatus.BIT_0,
		"B": gostatus.BIT_1,
		"C": gostatus.BIT_2,
	}

	st := gostatus.NewStatus(nil)
	fmt.Printf("[Initialized] bits: %s\n", st.GetBitsString())
	// [Initialized] bits: 00000000000000000000000000000000

	st.On(masks["A"])
	fmt.Printf("[Flag A is turned on] bits: %s\n", st.GetBitsString())
	// [Flag A is turned on] bits: 00000000000000000000000000000001

	st.On(masks["B"] | masks["C"])
	fmt.Printf("[Flag B,C are turned on] bits: %s\n", st.GetBitsString())
	// [Flag B,C are turned on] bits: 00000000000000000000000000000111

	st.Off(masks["A"] | masks["C"])
	fmt.Printf("[Flag A,C are turned off] bits: %s\n", st.GetBitsString())
	// [Flag A,C are turned off] bits: 00000000000000000000000000000010

	st.Off(masks["B"])
	fmt.Printf("[Flag B are turned off] bits: %s\n", st.GetBitsString())
	// [Flag B are turned off] bits: 00000000000000000000000000000000

	st2 := gostatus.NewStatus(&gostatus.Option{gostatus.GetAllOnBits()})
	fmt.Printf("[Bits are all turned on] bits: %s\n", st2.GetBitsString())
	// [Bits are all turned on] bits: 11111111111111111111111111111111

	st2.SetBits(gostatus.GetAllOffBits())
	fmt.Printf("[Bits are all turned off] bits: %s\n", st2.GetBitsString())
	// [Bits are all turned off] bits: 00000000000000000000000000000000

	st2.SetBitsString("00001111000011110000111100001111")
	fmt.Printf("[Set Bits by string bit sequence] bits: %s\n", st2.GetBitsString())
	// [Bits are all turned off] bits: 00001111000011110000111100001111

	// Get Bits as uint32
	st2.Off(gostatus.GetAllOnBits())
	st2.On(gostatus.BIT_3 | gostatus.BIT_2 | gostatus.BIT_1 | gostatus.BIT_0)
	fmt.Printf("[As uint32] bits: %d\n", st2.GetBitsUint32())
	// [As uint32] bits: 15
}
