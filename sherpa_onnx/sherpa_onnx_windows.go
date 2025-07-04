//go:build (windows && amd64) || (windows && 386)

package sherpa_onnx

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go-windows"
)

type OnlineTransducerModelConfig = sherpa.OnlineTransducerModelConfig
type FeatureConfig = sherpa.FeatureConfig
type HomophoneReplacerConfig = sherpa.HomophoneReplacerConfig
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
type OfflineDolphinModelConfig = sherpa.OfflineDolphinModelConfig
type OfflineZipformerCtcModelConfig = sherpa.OfflineZipformerCtcModelConfig
type OfflineWhisperModelConfig = sherpa.OfflineWhisperModelConfig
type OfflineFireRedAsrModelConfig = sherpa.OfflineFireRedAsrModelConfig
type OfflineTdnnModelConfig = sherpa.OfflineTdnnModelConfig
type OfflineSenseVoiceModelConfig = sherpa.OfflineSenseVoiceModelConfig
type OfflineMoonshineModelConfig = sherpa.OfflineMoonshineModelConfig
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
type OfflineTtsMatchaModelConfig = sherpa.OfflineTtsMatchaModelConfig
type OfflineTtsKokoroModelConfig = sherpa.OfflineTtsKokoroModelConfig
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

type OfflineSpeakerSegmentationPyannoteModelConfig = sherpa.OfflineSpeakerSegmentationPyannoteModelConfig
type OfflineSpeakerSegmentationModelConfig = sherpa.OfflineSpeakerSegmentationModelConfig
type FastClusteringConfig = sherpa.FastClusteringConfig
type OfflineSpeakerDiarizationConfig = sherpa.OfflineSpeakerDiarizationConfig
type OfflineSpeakerDiarization = sherpa.OfflineSpeakerDiarization
type OfflineSpeakerDiarizationSegment = sherpa.OfflineSpeakerDiarizationSegment

var NewOfflineSpeakerDiarization = sherpa.NewOfflineSpeakerDiarization
var DeleteOfflineSpeakerDiarization = sherpa.DeleteOfflineSpeakerDiarization

type OfflinePunctuationConfig = sherpa.OfflinePunctuationConfig
type OfflinePunctuationModelConfig = sherpa.OfflinePunctuationModelConfig
type OfflinePunctuation = sherpa.OfflinePunctuation

var NewOfflinePunctuation = sherpa.NewOfflinePunctuation
var DeleteOfflinePunc = sherpa.DeleteOfflinePunc

// ============================================================
// For Keyword spotter
// ============================================================

type KeywordSpotter = sherpa.KeywordSpotter
type KeywordSpotterConfig = sherpa.KeywordSpotterConfig
type KeywordSpotterResult = sherpa.KeywordSpotterResult

var NewKeywordSpotter = sherpa.NewKeywordSpotter
var NewKeywordStream = sherpa.NewKeywordStream
var NewKeywordStreamWithKeywords = sherpa.NewKeywordStreamWithKeywords
var DeleteKeywordSpotter = sherpa.DeleteKeywordSpotter

// ============================================================
// For Audio Tagging
// ============================================================

type AudioTagging = sherpa.AudioTagging
type AudioTaggingConfig = sherpa.AudioTaggingConfig
type AudioEvent = sherpa.AudioEvent

var NewAudioTagging = sherpa.NewAudioTagging
var NewAudioTaggingStream = sherpa.NewAudioTaggingStream
var DeleteAudioTagging = sherpa.DeleteAudioTagging

// ============================================================
// For Speech Enhancement
// ============================================================
type OfflineSpeechDenoiser = sherpa.OfflineSpeechDenoiser
type OfflineSpeechDenoiserConfig = sherpa.OfflineSpeechDenoiserConfig
type OfflineSpeechDenoiserModelConfig = sherpa.OfflineSpeechDenoiserModelConfig
type OfflineSpeechDenoiserGtcrnModelConfig = sherpa.OfflineSpeechDenoiserGtcrnModelConfig
type DenoisedAudio = sherpa.DenoisedAudio

var DeleteOfflineSpeechDenoiser = sherpa.DeleteOfflineSpeechDenoiser
var NewOfflineSpeechDenoiser = sherpa.NewOfflineSpeechDenoiser

var GetVersion = sherpa.GetVersion
var GetGitSha1 = sherpa.GetGitSha1
var GetGitDate = sherpa.GetGitDate
