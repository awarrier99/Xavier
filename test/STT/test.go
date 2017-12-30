package main

import (
	ts "github.com/awarrier99/Xavier/text_speech"
)

func main() {
	//ts.Recognize_File(os.Getenv("XAVIER") + "/audio/temp/temp_in.wav")
	ts.StreamingRecognize_Mic()
}
