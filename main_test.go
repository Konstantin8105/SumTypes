package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// unionValues - C code example:
// union {
// 		long i;
// 		long j;
// };
type unionValues struct {
	memory [8]byte
}

func TestValues(t *testing.T) {
	var uni unionValues

	var inp int64 = -1234567890
	uni.memory = *((*[8]byte)(unsafe.Pointer(&inp)))
	var res int64
	res = *((*int64)(unsafe.Pointer(&uni.memory)))

	// Output
	fmt.Printf("inp = %+v\n", inp)
	fmt.Printf("uni = %+v\n", uni)
	fmt.Printf("res = %+v\n", res)
	if res != inp {
		t.Fatalf("Result is not same. %v != %v", inp, res)
	}
}

// unionValuesAndArray - C code example:
// union {
// 		long i;
// 		int  j[2];
// };
type unionValuesAndArray struct {
	memory [8]byte
}

func TestValuesAndArray(t *testing.T) {
	var uni unionValuesAndArray

	var inp int64 = -1234567890
	uni.memory = *((*[8]byte)(unsafe.Pointer(&inp)))
	var ind [2]int32
	ind = *((*[2]int32)(unsafe.Pointer(&uni.memory)))
	var res int64
	res = *((*int64)(unsafe.Pointer(&ind)))

	// Output
	fmt.Printf("inp = %+v\n", inp)
	fmt.Printf("ind = %+v\n", ind)
	fmt.Printf("uni = %+v\n", uni)
	fmt.Printf("res = %+v\n", res)
	if res != inp {
		t.Fatalf("Result is not same. %v != %v", inp, res)
	}
}

// unionPointers - C code example:
// union {
// 		long * i;
// 		int  * j[2];
// };
type unionPointers struct {
	memory unsafe.Pointer
}

func TestPointers(t *testing.T) {
	var uni unionPointers

	var inp int64 = -1234567890
	uni.memory = unsafe.Pointer(&inp)
	var ind *[2]int32
	ind = (*[2]int32)(uni.memory)
	//ind[0] += 12222 // changes
	(*[2]int32)(uni.memory)[0] += 56980
	(*[2]int32)(uni.memory)[1] += 12222

	var res int64
	res = *((*int64)(unsafe.Pointer(ind)))

	// Output
	fmt.Printf("inp = %#v\n", inp)
	fmt.Printf("ind = %#v\n", ind)
	fmt.Printf("uni = %#v\n", uni)
	fmt.Printf("res = %#v\n", res)
	if res != inp {
		t.Fatalf("Result is not same. %v != %v", inp, res)
	}
}

type unionGo struct {
	memory unsafe.Pointer
}

var un unionGo

func GoInit() {
	var m [200]byte
	un.memory = unsafe.Pointer(&m)
	(*[2]int32)(un.memory)[0] = -1234
	(*[2]int32)(un.memory)[1] = 4556
}

func TestCtoGo(t *testing.T) {
	CInit()
	GoInit()

	fmt.Println("0: ", GetI0(), (*[2]int32)(un.memory)[0]) // int
	fmt.Println("1: ", GetI1(), (*[2]int32)(un.memory)[1]) // int
	fmt.Println("2: ", GetL(), *(*int64)(un.memory))       // long
	fmt.Println("3: ", GetC(), *(*[8]uint16)(un.memory))   // char[8]
	fmt.Println("4: ", GetI(), *(*[2]int32)(un.memory))    // int[2]
	fmt.Println("5: ", GetSh(), *(*[16]uint8)(un.memory))  // short[4]
	fmt.Println("6: ", GetD(), *(*float64)(un.memory))     // double
}
