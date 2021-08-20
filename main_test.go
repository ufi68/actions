package main

import (
    "testing"
)

func TestFoo(t *testing.T) {
  if foo("Hello Word") != "Hello World" { t.Error(foo()) }
}
