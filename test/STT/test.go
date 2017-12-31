package main

import (
	ts "github.com/awarrier99/Xavier/text_speech"
)

func main() {
	//ts.Recognize_File(os.Getenv("XAVIER") + "/audio/test/test_record.wav")
	ts.StreamingRecognize_Mic()
}
