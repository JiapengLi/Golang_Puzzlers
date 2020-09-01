package lib

import (
	"os"
	in "puzzlers/demo06/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
