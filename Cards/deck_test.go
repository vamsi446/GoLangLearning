package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != 16{
		t.Errorf("Expected deck of length 20, but got %v",len(d))
	}
}