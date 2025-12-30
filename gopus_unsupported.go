//go:build !windows
// +build !windows

package gopus

import (
	"errors"
	"fmt"
)

type Encoder uintptr // opaque pointer to an Opus encoder

type Application int

const (
	// Opus Application Constants
	OPUS_APPLICATION_VOIP                Application = 2048
	OPUS_APPLICATION_AUDIO               Application = 2049
	OPUS_APPLICATION_RESTRICTED_LOWDELAY Application = 2051

	// Opus Error Codes
	OPUS_OK               = 0
	OPUS_BAD_ARG          = -1
	OPUS_BUFFER_TOO_SMALL = -2
	OPUS_INTERNAL_ERROR   = -3
	OPUS_INVALID_PACKET   = -4
	OPUS_UNIMPLEMENTED    = -5
	OPUS_INVALID_STATE    = -6
	OPUS_ALLOC_FAIL       = -7
)

type Error int32

func (e Error) Error() string {
	return fmt.Sprintf("opus: error code %d", int32(e))
}

func CreateEncoder(fs int, channels int, application Application) (Encoder, error) {
	return 0, errors.New("gopus: not implemented on this platform")
}

func (e Encoder) Destroy() {}
