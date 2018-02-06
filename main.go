package main

/*
#include "stdio.h"

union c_union{
	int    i[2];	// 2 * 32
	short  sh[4];	// 4 * 16
	long   l   ;	// 1 * 64
	char   c[80];	// 8 *  8
	double d    ;	// 1 * 64
};

union c_union u;

void init()
{
	u.i[0] = -1234;
	u.i[1] = 4556;
}

int * getI()
{
	return u.i;
}

short * getSh()
{
	return u.sh;
}

int getI0()
{
	return u.i[0];
}

int getI1()
{
	return u.i[1];
}

long getL()
{
	return u.l;
}

char * getC()
{
	return u.c;
}

double getD()
{
	return u.d;
}
*/
import "C"
import "unsafe"

func main() {
}

func CInit() {
	C.init()
}

func GetI() [2]int32 {
	return *((*[2]int32)(unsafe.Pointer(C.getI())))
}

func GetSh() [16]uint8 {
	return *((*[16]uint8)(unsafe.Pointer(C.getSh())))
}

func GetI0() int32 {
	return int32(C.getI0())
}

func GetI1() int32 {
	return int32(C.getI1())
}

func GetL() int64 {
	return int64(C.getL())
}

func GetC() [8]uint16 {
	return *((*[8]uint16)(unsafe.Pointer(C.getC())))
}

func GetD() float64 {
	return float64(C.getD())
}
