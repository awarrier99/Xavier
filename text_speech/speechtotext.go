package text_speech

import (
    "fmt"
    "io/ioutil"
    "log"

    "golang.org/x/net/context"

    speech "cloud.google.com/go/speech/apiv1"
    speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func STT() {
    ctx := context.Background()

    client, err := speech.NewClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    filename := "../2017-12-21-14:18:25.wav"

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }

    resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
        Config: &speechpb.RecognitionConfig{
            Encoding: speechpb.RecognitionConfig_LINEAR16,
            LanguageCode: "en-US",
        },
        Audio: &speechpb.RecognitionAudio{
            AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
        },
    })
    if err != nil {
        log.Fatalf("Failed to recognize: %v", err)
    }

    for _, result := range resp.Results {
        for _, alt := range result.Alternatives {
            fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
        }
    }
    fmt.Println("Done")
}
