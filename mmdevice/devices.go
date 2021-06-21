package mmdevice

// +build !linux

//#include "devices.h"
import "C"

func Hej() {
	C.hej();
}

func Cleanup() {
	C.cleanup();
}