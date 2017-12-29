package main

import (
	"os"

	ts "github.com/awarrier99/Xavier/text_speech"
)

func main() {
	ts.Recognize(os.Getenv("XAVIER") + "/test_record.wav")
}
