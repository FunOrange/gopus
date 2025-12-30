package gopus

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	opus                 = windows.NewLazyDLL("opus.dll")
	opus_encoder_create  = opus.NewProc("opus_encoder_create")
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

func (e Encoder) Destroy() {
	if e != 0 {
		opus_encoder_destroy.Call(uintptr(e))
	}
}
