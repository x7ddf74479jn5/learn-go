package main

import (
	"os"
	"time"

	"github.com/x7ddf74479jn5/learn-go/with-tests/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
