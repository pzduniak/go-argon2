package argon2

// #cgo CFLAGS: -I${SRCDIR}/phc-winner-argon2/include
// #include <stdlib.h>
// #include "phc-winner-argon2/src/argon2.c"
// #include "phc-winner-argon2/src/core.c"
// #include "phc-winner-argon2/src/blake2/blake2b.c"
// #include "phc-winner-argon2/src/thread.c"
// #include "phc-winner-argon2/src/encoding.c"
// #include "phc-winner-argon2/src/opt.c"
import "C"
