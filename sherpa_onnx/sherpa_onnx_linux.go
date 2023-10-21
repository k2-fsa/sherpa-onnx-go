//go:build (!android && linux && arm64) || (!android && linux && amd64 && !musl) || (!android && linux && arm && !arm7) || (!android && arm7) || (!android && linux && 386 && !musl) || (!android && musl) || (!android && linux && mips) || (!android && linux && mips64) || (!android && linux && mips64le) || (!android && linux && mipsle)

package sherpa_onnx

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go-linux"
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

// ============================================================
// For offline ASR (i.e., non-streaming ASR)
// ============================================================
type OfflineTransducerModelConfig = sherpa.OfflineTransducerModelConfig
type OfflineParaformerModelConfig = sherpa.OfflineParaformerModelConfig
type OfflineNemoEncDecCtcModelConfig = sherpa.OfflineNemoEncDecCtcModelConfig
type OfflineLMConfig = sherpa.OfflineLMConfig
type OfflineModelConfig = sherpa.OfflineModelConfig
type OfflineRecognizerConfig = sherpa.OfflineRecognizerConfig
type OfflineRecognizer = sherpa.OfflineRecognizer
type OfflineStream = sherpa.OfflineStream
type OfflineRecognizerResult = sherpa.OfflineRecognizerResult

var NewOfflineRecognizer = sherpa.NewOfflineRecognizer
var DeleteOfflineRecognizer = sherpa.DeleteOfflineRecognizer

var NewOfflineStream = sherpa.NewOfflineStream
var DeleteOfflineStream = sherpa.DeleteOfflineStream

type OfflineTtsVitsModelConfig = sherpa.OfflineTtsVitsModelConfig
type OfflineTtsModelConfig = sherpa.OfflineTtsModelConfig
type OfflineTtsConfig = sherpa.OfflineTtsConfig
type GeneratedAudio = sherpa.GeneratedAudio
type OfflineTts = sherpa.OfflineTts

var DeleteOfflineTts = sherpa.DeleteOfflineTts
var NewOfflineTts = sherpa.NewOfflineTts
