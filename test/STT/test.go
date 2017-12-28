package main

import ts "github.com/awarrier99/Xavier/text_speech"

func main() {
    ts.STT()
    ts.Say("Hello")
    ts.Say("Guten Tag", "de")
}
