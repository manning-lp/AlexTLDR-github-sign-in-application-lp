package main

import (
	"fmt"
	"github.com/username/hello-world/hello"
	"github.com/username/hello-world/world"
	"io"
	"os"
)

func main() {
	displayGreetings(os.Stdout)
}

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}
