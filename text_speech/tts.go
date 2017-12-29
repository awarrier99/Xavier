package text_speech

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/awarrier99/Xavier/errcheck"
	_ "github.com/faiface/beep"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

const (
	aKeyStr = "AWSAccessKeyId="
	sKeyStr = "AWSSecretKey="
)

func getKeys(filename string) (string, string, error) {
	f, err := ioutil.ReadFile(filename)
	errcheck.Err(err)

	contents := string(f)

	aPos := strings.Index(contents, aKeyStr) + len([]rune(aKeyStr))
	aPosEnd := strings.Index(contents, "\n")
	sPos := strings.Index(contents, sKeyStr) + len([]rune(sKeyStr))

	if (aPos == -1) || (sPos == -1) {
		return "", "", errors.New("Keys not found in file")
	}

	aKey := contents[aPos:aPosEnd]
	sKey := contents[sPos:]

	return aKey, sKey, nil
}

func Say(s string) io.ReadCloser {
	// aKey, sKey, err := getKeys("/Users/Ashvin/go/src/github.com/awarrier99/Xavier/keys/rootkey.csv")
	// errcheck.Err(err)

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
	return stream
}
