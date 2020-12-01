# Status management with bits.


### Installation
```
$ go get github.com/kentaro-a/gostatus
```


### How to use 
Set 32 bits mask gostatus.Bits as a status list and pass it to gostatus.NewStatus().  
You can turn On/Off any bits using gostatus.Bits.  

```
masks := map[string]Bits{
"A": BIT_0,
"B": BIT_1,
"C": BIT_2,
}

st := NewStatus(nil)
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

```
