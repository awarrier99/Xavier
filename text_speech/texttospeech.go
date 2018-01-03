package text_speech

import (
	"io"
	"os"

	"github.com/awarrier99/Xavier/errcheck"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

func cache(s string) {
	//TODO
}

func getFromCache(s string) (string, error) {
	//TODO
	return "", nil
}

func split(r rune) bool {
	return r == '.' || r == '?' || r == '!'
}

func Say(s string) {
	// sentences := strings.FieldsFunc(s, split)
	//
	// for i := 0; i < len(sentences); i++ {
	// 	sentences[i] = strings.TrimSpace(sentences[i])
	// 	fmt.Println(sentences[i])
	// }

	s += " a" //Audio cuts out early so this addition allows user-specified sentence to be said

	svc := polly.New(session.New(), &aws.Config{
		Credentials: credentials.NewEnvCredentials(),
		Region:      aws.String("us-east-1"),
	})
	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		SampleRate:   aws.String("22050"),
		Text:         aws.String(s),
		TextType:     aws.String("text"),
		VoiceId:      aws.String("Brian"),
	}

	result, err := svc.SynthesizeSpeech(input)
	errcheck.Err(err)
	stream := result.AudioStream

	f, err := os.Create(os.Getenv("XAVIER") + "/audio/temp/temp_out.mp3")
	errcheck.Err(err)
	_, err = io.Copy(f, stream)
	errcheck.Err(err)

	play(os.Getenv("XAVIER") + "/audio/temp/temp_out.mp3")
}
