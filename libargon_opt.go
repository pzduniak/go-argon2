// +build amd64 386

package argon2

// #cgo CFLAGS: -I${SRCDIR}/libargon2/include
// #include <stdlib.h>
// #include "libargon2/src/argon2.c"
// #include "libargon2/src/core.c"
// #include "libargon2/src/blake2/blake2b.c"
// #include "libargon2/src/thread.c"
// #include "libargon2/src/encoding.c"
// #include "libargon2/src/opt.c"
import "C"
