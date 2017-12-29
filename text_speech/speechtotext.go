package text_speech

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"

	"github.com/awarrier99/Xavier/errcheck"
)

func Recognize(filename string) {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	errcheck.Err(err)

	data, err := ioutil.ReadFile(filename)
	errcheck.Err(err)

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:     speechpb.RecognitionConfig_LINEAR16,
			LanguageCode: "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	errcheck.Err(err)

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	fmt.Println("Done")
}
