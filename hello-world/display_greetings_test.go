package main

import (
	"bytes"
	"testing"
)

func TestDisplayGreetings(t *testing.T) {
	buf := new(bytes.Buffer)

	displayGreetings(buf)
	expectedOutput := "Hello!\nWorld!\n"
	gotOutput := buf.String()
	if gotOutput != expectedOutput {
		t.Errorf("Expected output to be %q but got %q", expectedOutput, gotOutput)
	}
}
