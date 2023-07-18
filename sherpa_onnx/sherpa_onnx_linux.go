//go:build (!android && linux && arm64) || (!android && linux && amd64 && !musl) || (!android && linux && arm && !arm7) || (!android && arm7) || (!android && linux && 386 && !musl) || (!android && musl) || (!android && linux && mips) || (!android && linux && mips64) || (!android && linux && mips64le) || (!android && linux && mipsle)

package sherpa_onnx

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go-linux/sherpa-onnx"
)

type OnlineTransducerModelConfig = sherpa.OnlineTransducerModelConfig
type FeatureConfig = sherpa.FeatureConfig
type OnlineRecognizerConfig = sherpa.OnlineRecognizerConfig
type OnlineRecognizerResult = sherpa.OnlineRecognizerResult
type OnlineRecognizer = sherpa.OnlineRecognizer
type OnlineStream = sherpa.OnlineStream

var NewOnlineRecognizer = sherpa.NewOnlineRecognizer
var DeleteOnlineRecognizer = sherpa.DeleteOnlineRecognizer

var NewOnlineStream = sherpa.NewOnlineStream
var DeleteOnlineStream = sherpa.DeleteOnlineStream
