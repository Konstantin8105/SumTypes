package main_test

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

// unionValuesByInterface - C code example:
// union {
// 		long i;
// };
type unionValuesByInterface struct {
	memory interface{}
}

func TestValuesByInterface(t *testing.T) {
	var uni unionValuesByInterface

	var inp int64 = -1234567890
	uni.memory = (interface{})(inp)
	var res int64
	res = (*((*interface{})(unsafe.Pointer(&uni.memory)))).(int64)

	// Output
	fmt.Printf("inp = %+v\n", inp)
	fmt.Printf("uni = %+v\n", uni)
	fmt.Printf("res = %+v\n", res)
	if res != inp {
		t.Fatalf("Result is not same. %v != %v", inp, res)
	}
}

// unionWithArray - C code example:
// union {
// 	int i[2];
// };
type unionWithArray struct {
	memory [8]byte
}

func TestArray(t *testing.T) {
	var uni unionWithArray

	inp := [2]int32{-1234567, 56}
	uni.memory = *((*[8]byte)(unsafe.Pointer(&inp)))
	var res [2]int32
	res = *((*[2]int32)(unsafe.Pointer(&uni.memory)))

	// Output
	fmt.Printf("inp = %+v\n", inp)
	fmt.Printf("uni = %+v\n", uni)
	fmt.Printf("res = %+v\n", res)
	if res != inp {
		t.Fatalf("Result is not same. %v != %v", inp, res)
	}
}
