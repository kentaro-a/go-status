package main

import (
	"fmt"

	"github.com/kentaro-a/gostatus"
)

func main() {

	masks := map[string]gostatus.BitMask{
		"A": gostatus.BIT_0,
		"B": gostatus.BIT_1,
		"C": gostatus.BIT_2,
	}

	st := gostatus.NewStatus(nil)
	fmt.Printf("[Initialized] bits: %s\n", st.Bits())
	// [Initialized] bits: 00000000

	st.On(masks["A"])
	fmt.Printf("[Flag A is turned on] bits: %s\n", st.Bits())
	// [Flag A is turned on] bits: 00000001

	st.On(masks["B"] | masks["C"])
	fmt.Printf("[Flag B,C are turned on] bits: %s\n", st.Bits())
	// [Flag B,C are turned on] bits: 00000111

	st.Off(masks["A"] | masks["C"])
	fmt.Printf("[Flag A,C are turned off] bits: %s\n", st.Bits())
	// [Flag A,C are turned off] bits: 00000010

	st.Off(masks["B"])
	fmt.Printf("[Flag B are turned off] bits: %s\n", st.Bits())
	// [Flag B are turned off] bits: 00000000
}
