package main

import (
    "testing"
)

func TestSayWelcome(t *testing.T) {
    if SayWelcome() != "Hi there! Welcome" {
        t.Error("Greetings TestCase failed ::: Expected Result ::: Hi there! Welcome")
    }
}
