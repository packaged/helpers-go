package helpers

import (
	"testing"
)

func TestFormUrlDecode(t *testing.T) {
	s1 := urlDecodeSchema1{}
	err := FormURLDecode(&s1, "one=1&Two=2")
	if err != nil {
		t.Fatalf("Error decoding: %s", err)
	}
	if s1.One != "1" {
		t.Fatalf("Expected One to be 1, got %s", s1.One)
	}
	if s1.Two != "2" {
		t.Fatalf("Expected Two to be 2, got %s", s1.Two)
	}
}

func TestFormUrlDecodeFailing(t *testing.T) {
	s1 := urlDecodeSchema1{}
	err := FormURLDecode(&s1, "%")
	if err == nil {
		t.Fatal("Expected error decoding, got nil")
	}
}

type urlDecodeSchema1 struct {
	One string `schema:"one"`
	Two string
}
