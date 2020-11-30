package main

import (
	"fmt"

	"github.com/kentaro-a/gostatus"
)

func main() {

	masks := map[string]gostatus.BitMask{
		"A": 0b00000001,
		"B": 0b00000010,
		"C": 0b00000100,
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
