//go:build windows
// +build windows

package gopus

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func TestCreateAndDestroyEncoder(t *testing.T) {
	// create encoder
	fs := 48000
	channels := 2
	application := OPUS_APPLICATION_AUDIO
	encoder, err := CreateEncoder(fs, channels, application)
	if err != nil {
		t.Fatal(err)
	}

	// destroy encoder
	encoder.Destroy()
}

func TestCreateEncoderWithBadArguments(t *testing.T) {
	fs := 48000
	channels := -1
	application := OPUS_APPLICATION_AUDIO
	e, err := CreateEncoder(fs, channels, application)
	if e != 0 {
		t.Fatal("expected encoder to be 0")
	}
	if err != nil {
		var oErr Error
		if errors.As(err, &oErr) {
			switch oErr {
			case OPUS_BAD_ARG:
				return
			default:
				t.Fatal("expected bad argument error")
			}
		}
	}
	t.Fatal("expected bad argument error")
}

func TestEncodeSilence(t *testing.T) {
	fs := 48000
	channels := 1
	application := OPUS_APPLICATION_AUDIO
	encoder, err := CreateEncoder(fs, channels, application)
	if err != nil {
		t.Fatal(err)
	}
	defer encoder.Destroy()

	frameSize := int32(960)
	pcm := make([]int16, frameSize)
	out := make([]byte, frameSize)
	ret, err := encoder.EncodeInt16(pcm, frameSize, out)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("encoded %d bytes -> %d bytes", len(pcm)*2, ret)

	// TODO: idk how to verify encoded data

	fmt.Printf("\n")
}

func TestEncodeNoise(t *testing.T) {
	fs := 48000
	channels := 1
	application := OPUS_APPLICATION_AUDIO
	encoder, err := CreateEncoder(fs, channels, application)
	if err != nil {
		t.Fatal(err)
	}
	defer encoder.Destroy()

	frameSize := int32(960)
	pcm := make([]int16, frameSize)
	for i := range pcm {
		pcm[i] = int16(rand.Intn(1000) - 500)
	}
	out := make([]byte, frameSize)
	ret, err := encoder.EncodeInt16(pcm, frameSize, out)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("encoded %d bytes -> %d bytes", len(pcm)*2, ret)
	fmt.Printf("\n")
}

func TestEncodeFloat(t *testing.T) {
	fs := 48000
	channels := 1
	application := OPUS_APPLICATION_AUDIO
	encoder, err := CreateEncoder(fs, channels, application)
	if err != nil {
		t.Fatal(err)
	}
	defer encoder.Destroy()

	frameSize := int32(960)
	pcm := make([]float32, frameSize)
	for i := range pcm {
		pcm[i] = float32((2 * rand.Float32()) - 1)
	}
	out := make([]byte, frameSize)
	ret, err := encoder.EncodeFloat32(pcm, frameSize, out)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("encoded %d bytes -> %d bytes", len(pcm)*4, ret)
	fmt.Printf("\n")
}
