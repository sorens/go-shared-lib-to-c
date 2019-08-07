## Go Shared Library To C

Information gleaned from
- https://github.com/vladimirvivien/go-cshared-examples
- http://geekwentfreak.com/posts/golang/cgo_pass_receive_string_c/
- http://blog.ralch.com/tutorial/golang-sharing-libraries/
- https://dave.cheney.net/tag/cgo
- https://golang.org/cmd/cgo/

1. Create Go code that does something

```go
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
```

2. Create the c code uses the Go code

```c
#include <stdio.h>

#include "l.h"

int main() 
{
    printf("A simple c program that calls Go code...\n");
    GoString world = {"world", 5};
    printf(" => %s\n", Hello(world));
    GoString nurse = {"nurse", 5};
    printf(" => %s\n", Hello(nurse));
}
```

3. build the go shared library and create the header

```shell
go build -o l.so -buildmode=c-shared l.go
```

4. Compile the c code with the Go shared library

```shell
gcc -o l l.c ./l.so
```

5. Run the application

```shell
> ./l
A simple c program that calls Go code...
GOLOG: hello, world
 => hello, world
GOLOG: hello, nurse
 => hello, nurse
```

