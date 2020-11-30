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
	fmt.Println(st.Bits())
	fmt.Println(st.IsOn(masks["A"]))
	fmt.Println(st.IsOn(masks["B"]))
	fmt.Println(st.IsOn(masks["C"]))

	st.On(masks["A"])

	fmt.Println(st.Bits())
	fmt.Println(st.IsOn(masks["A"]))
	fmt.Println(st.IsOn(masks["B"]))
	fmt.Println(st.IsOn(masks["C"]))
}
