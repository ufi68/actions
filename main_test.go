package main

import (
    "testing"
)

func TestFoo(t *testing.T) {
  if foo() != "Hello World" { t.Error(foo()) }
}
