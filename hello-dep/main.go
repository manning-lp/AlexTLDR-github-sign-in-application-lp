package main

import (
	"fmt"
	"github.com/amitsaha/using-go-modules/greetings/hello"
	"github.com/amitsaha/using-go-modules/greetings/world"
	"io"
	"os"
)

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}

func main() {
	displayGreetings(os.Stdout)
}
