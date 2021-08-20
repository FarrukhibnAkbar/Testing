package main

import (
	"App/calc"
	"testing"
)

func TestAdd(t *testing.T) {
  got := basic.Add(5,6)
  if got != 11 {
    t.Errorf("Error")
  } 
 
}