//go:build windows
// +build windows

package gopus

import (
	"errors"
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
	_, err := CreateEncoder(fs, channels, application)
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
