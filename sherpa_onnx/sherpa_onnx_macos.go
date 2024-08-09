//go:build (darwin && amd64 && !ios) || (darwin && arm64 && !ios)

package sherpa_onnx

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go-macos"
)

type OnlineTransducerModelConfig = sherpa.OnlineTransducerModelConfig
type FeatureConfig = sherpa.FeatureConfig
type OnlineCtcFstDecoderConfig = sherpa.OnlineCtcFstDecoderConfig
type OnlineRecognizerConfig = sherpa.OnlineRecognizerConfig
type OnlineRecognizerResult = sherpa.OnlineRecognizerResult
type OnlineRecognizer = sherpa.OnlineRecognizer
type OnlineStream = sherpa.OnlineStream
type OnlineModelConfig = sherpa.OnlineModelConfig

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
type Wave = sherpa.Wave
type OfflineTts = sherpa.OfflineTts

var DeleteOfflineTts = sherpa.DeleteOfflineTts
var NewOfflineTts = sherpa.NewOfflineTts

type SileroVadModelConfig = sherpa.SileroVadModelConfig
type VadModelConfig = sherpa.VadModelConfig
type CircularBuffer = sherpa.CircularBuffer

var DeleteCircularBuffer = sherpa.DeleteCircularBuffer
var NewCircularBuffer = sherpa.NewCircularBuffer

type SpeechSegment = sherpa.SpeechSegment
type VoiceActivityDetector = sherpa.VoiceActivityDetector

var NewVoiceActivityDetector = sherpa.NewVoiceActivityDetector
var DeleteVoiceActivityDetector = sherpa.DeleteVoiceActivityDetector

type SpokenLanguageIdentificationWhisperConfig = sherpa.SpokenLanguageIdentificationWhisperConfig
type SpokenLanguageIdentificationConfig = sherpa.SpokenLanguageIdentificationConfig
type SpokenLanguageIdentification = sherpa.SpokenLanguageIdentification
type SpokenLanguageIdentificationResult = sherpa.SpokenLanguageIdentificationResult

var NewSpokenLanguageIdentification = sherpa.NewSpokenLanguageIdentification
var DeleteSpokenLanguageIdentification = sherpa.DeleteSpokenLanguageIdentification

type SpeakerEmbeddingExtractorConfig = sherpa.SpeakerEmbeddingExtractorConfig
type SpeakerEmbeddingExtractor = sherpa.SpeakerEmbeddingExtractor

var NewSpeakerEmbeddingExtractor = sherpa.NewSpeakerEmbeddingExtractor
var DeleteSpeakerEmbeddingExtractor = sherpa.DeleteSpeakerEmbeddingExtractor

type SpeakerEmbeddingManager = sherpa.SpeakerEmbeddingManager

var NewSpeakerEmbeddingManager = sherpa.NewSpeakerEmbeddingManager
var DeleteSpeakerEmbeddingManager = sherpa.DeleteSpeakerEmbeddingManager
var ReadWave = sherpa.ReadWave
