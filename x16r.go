package x16r

/*
void x16r_hash(const char* input, char* output);
#cgo LDFLAGS: -L. -Llib -lx16r
 */
import "C"
import (
	"unsafe"
)

func X16R(data []byte) []byte {
	output := make([]byte, 32)
	C.x16r_hash((*C.char)(unsafe.Pointer(&data[0])), (*C.char)(unsafe.Pointer(&output[0])))
	return output
}
