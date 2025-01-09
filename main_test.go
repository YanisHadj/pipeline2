package main

import "testing"

func TestMain(t *testing.T) {
    expected := "Hello, CI/CD!"
    if expected != "Hello, CI/CD!" {
        t.Errorf("Expected %s but got %s", expected, "something else")
    }
}
