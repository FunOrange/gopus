## Go wrapper for Opus

This package provides Go bindings for the xiph.org C Opus encoder (opus.dll)

github.com/hraban/opus wasn't working for me and I only needed an encoder, so I wrote this.

Features:

- ✅ encode raw PCM data to raw Opus data (not .ogg, the raw frames defined in [RFC7845](https://datatracker.ietf.org/doc/html/rfc6716#section-3]))
- ✅ no CGO (uses syscall to dynamically link to opus.dll)
- ❌ only exposes 3 functions from opus.h for now (the minimum required to encode something)
- ❌ no decoding
- ❌ Windows only for now

## Requirements

You must acquire `opus.dll` and put it in the same folder as your built `.exe`.
You can get `opus.dll` by building from source obtained from https://opus-codec.org/downloads/.
Alternatively, you can download a pre-built `opus.dll` [here](https://www.youtube.com/watch?v=dQw4w9WgXcQ).

### Importing the library

```go
import "github.com/FunOrange/gopus"
```

### Usage

This example creates an opus encoder, encodes a 20ms frame of silence, then destroys the encoder.

```go
func main() {
	fs := 48000
	channels := 1
	application := gopus.OPUS_APPLICATION_AUDIO
	encoder, err := gopus.CreateEncoder(fs, channels, application)
	if err != nil {
		t.Fatal(err)
	}
	defer encoder.Destroy()

	frameSize := int32(960)
	pcm := make([]int16, frameSize)
	out := make([]byte, frameSize)
	ret, err := encoder.EncodeInt16(pcm, frameSize, out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("encoded %d bytes -> %d bytes\n", len(pcm)*2, ret)
}
```
