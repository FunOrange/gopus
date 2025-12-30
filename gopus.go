//go:build windows
// +build windows

package gopus

import (
	"errors"
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	opus                 = windows.NewLazyDLL("opus.dll")
	opus_encoder_create  = opus.NewProc("opus_encoder_create")
	opus_encode          = opus.NewProc("opus_encode")
	opus_encode_float    = opus.NewProc("opus_encode_float")
	opus_encoder_destroy = opus.NewProc("opus_encoder_destroy")
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
	var errCode int32
	encoder, _, _ := opus_encoder_create.Call(
		uintptr(int32(fs)),
		uintptr(int32(channels)),
		uintptr(int32(application)),
		uintptr(unsafe.Pointer(&errCode)),
	)
	if errCode != OPUS_OK {
		return 0, Error(errCode)
	}
	return Encoder(encoder), nil
}

// Encode encodes an Opus frame.
// pcm: Input signal (interleaved if 2 channels).
// frameSize: Number of samples per frame of input signal (e.g., 960 for 20ms at 48kHz).
// maxDataBytes: Size of the output buffer.
// Returns the length of the encoded packet in bytes.
func (e Encoder) EncodeInt16(pcm []int16, frameSize int32, data []byte) (int32, error) {
	if e == 0 {
		return 0, errors.New("encoder is not initialized")
	}
	if len(pcm) == 0 {
		return 0, errors.New("pcm is empty")
	}

	// We pass the address of the first element of the slices.
	// We cast the length of 'data' to int32 for the max_data_bytes parameter.
	ret, _, _ := opus_encode.Call(
		uintptr(e),
		uintptr(unsafe.Pointer(&pcm[0])),
		uintptr(frameSize),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(int32(len(data))),
	)

	// In opus_encode, a negative return value is an error code.
	// A positive value is the number of bytes written to 'data'.
	res := int32(ret)
	if res < 0 {
		return 0, Error(res)
	}

	return res, nil
}

func (e Encoder) EncodeFloat32(pcm []float32, frameSize int32, data []byte) (int32, error) {
	if e == 0 {
		return 0, errors.New("encoder is not initialized")
	}
	if len(pcm) == 0 {
		return 0, errors.New("pcm is empty")
	}

	// We pass the address of the first element of the slices.
	// We cast the length of 'data' to int32 for the max_data_bytes parameter.
	ret, _, _ := opus_encode_float.Call(
		uintptr(e),
		uintptr(unsafe.Pointer(&pcm[0])),
		uintptr(frameSize),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(int32(len(data))),
	)

	// In opus_encode, a negative return value is an error code.
	// A positive value is the number of bytes written to 'data'.
	res := int32(ret)
	if res < 0 {
		return 0, Error(res)
	}

	return res, nil
}

// NOTE: DO NOT CALL THIS TWICE!
func (e Encoder) Destroy() {
	if e != 0 {
		opus_encoder_destroy.Call(uintptr(e))
	}
}
