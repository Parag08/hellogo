package main

import (
    "testing"
)

func TestPrintHelloWorld(t *testing.T) {
    if PrintHelloWorld() != "Hello, world.\n" {
        t.Error("expected hello world") 
    }
}
