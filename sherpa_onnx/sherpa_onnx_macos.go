//go:build (darwin && amd64 && !ios) || (darwin && arm64 && !ios)

package sherpa_onnx

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go-macos/sherpa-onnx"
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
