package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func SpeechToText(filename string) {
	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx, option.WithCredentialsFile("cred.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println(filename, "| data len:", len(data))
	// Detects speech in the audio file.

	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		panic(err)
	}

	stream.Send(&speechpb.StreamingRecognizeRequest{
		// StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
			
		// },
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					
				},
			},
		},
	})

	
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:          speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz:   24000,
			AudioChannelCount: 1,
			LanguageCode:      "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			// AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
			AudioSource: &speechpb.RecognitionAudio_Content{},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	// Prints the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}
