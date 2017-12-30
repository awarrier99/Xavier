package text_speech

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/context"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"

	"github.com/awarrier99/Xavier/errcheck"
)

func StreamingRecognize_Mic() {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	errcheck.Err(err)

	stream, err := client.StreamingRecognize(ctx)
	errcheck.Err(err)

	err = stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 16000,
					LanguageCode:    "en-US",
				},
			},
		},
	})
	errcheck.Err(err)

	go func() {
		buf := make([]byte, 64)
		for {
			n, err := os.Stdin.Read(buf)
			if err == io.EOF {
				err = stream.CloseSend()
				errcheck.Err(err)
				return
			}

			if err != nil {
				continue
			}

			err = stream.Send(&speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
					AudioContent: buf[:n],
				},
			})
			errcheck.Err(err)
		}
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		errcheck.Err(err)

		for _, result := range resp.Results {
			fmt.Printf("Result: %+v\n", result)
		}
	}
}

func Recognize_File(filename string) {
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
