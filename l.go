package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"
import "sync"
import "unsafe"

var mu sync.Mutex

//export Hello
func Hello(message string) *C.char {
	mu.Lock()
	defer mu.Unlock()

	s := fmt.Sprintf("hello, %s", message)
	fmt.Println("GOLOG:", s)
	c_string := C.CString(s)
	defer C.free(unsafe.Pointer(c_string))
	return c_string
}

func main() {}
