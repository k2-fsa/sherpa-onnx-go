//go:build (darwin && amd64 && !ios) || (darwin && arm64 && !ios)
package sherpa_onnx

// ============================================================
// Code Generated Automatically for macos platform, DO NOT EDIT MANUALLY!!
// ============================================================

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go-macos"
)

// ============================================================
// Struct/Function Defines
// ============================================================

type OnlineTransducerModelConfig = sherpa.OnlineTransducerModelConfig
type OnlineParaformerModelConfig = sherpa.OnlineParaformerModelConfig
type OnlineZipformer2CtcModelConfig = sherpa.OnlineZipformer2CtcModelConfig
type OnlineNemoCtcModelConfig = sherpa.OnlineNemoCtcModelConfig
type OnlineToneCtcModelConfig = sherpa.OnlineToneCtcModelConfig
type OnlineModelConfig = sherpa.OnlineModelConfig
type FeatureConfig = sherpa.FeatureConfig
type OnlineCtcFstDecoderConfig = sherpa.OnlineCtcFstDecoderConfig
type HomophoneReplacerConfig = sherpa.HomophoneReplacerConfig
type OnlineRecognizerConfig = sherpa.OnlineRecognizerConfig
type OnlineRecognizerResult = sherpa.OnlineRecognizerResult
type OnlineRecognizer = sherpa.OnlineRecognizer
type OnlineStream = sherpa.OnlineStream
type OfflineTransducerModelConfig = sherpa.OfflineTransducerModelConfig
type OfflineParaformerModelConfig = sherpa.OfflineParaformerModelConfig
type OfflineNemoEncDecCtcModelConfig = sherpa.OfflineNemoEncDecCtcModelConfig
type OfflineZipformerCtcModelConfig = sherpa.OfflineZipformerCtcModelConfig
type OfflineWenetCtcModelConfig = sherpa.OfflineWenetCtcModelConfig
type OfflineOmnilingualAsrCtcModelConfig = sherpa.OfflineOmnilingualAsrCtcModelConfig
type OfflineMedAsrCtcModelConfig = sherpa.OfflineMedAsrCtcModelConfig
type OfflineFireRedAsrCtcModelConfig = sherpa.OfflineFireRedAsrCtcModelConfig
type OfflineDolphinModelConfig = sherpa.OfflineDolphinModelConfig
type OfflineWhisperModelConfig = sherpa.OfflineWhisperModelConfig
type OfflineCanaryModelConfig = sherpa.OfflineCanaryModelConfig
type OfflineFireRedAsrModelConfig = sherpa.OfflineFireRedAsrModelConfig
type OfflineFunASRNanoModelConfig = sherpa.OfflineFunASRNanoModelConfig
type OfflineMoonshineModelConfig = sherpa.OfflineMoonshineModelConfig
type OfflineTdnnModelConfig = sherpa.OfflineTdnnModelConfig
type OfflineSenseVoiceModelConfig = sherpa.OfflineSenseVoiceModelConfig
type OfflineLMConfig = sherpa.OfflineLMConfig
type OfflineModelConfig = sherpa.OfflineModelConfig
type OfflineRecognizerConfig = sherpa.OfflineRecognizerConfig
type OfflineRecognizer = sherpa.OfflineRecognizer
type OfflineStream = sherpa.OfflineStream
type OfflineRecognizerResult = sherpa.OfflineRecognizerResult
type OfflineTtsVitsModelConfig = sherpa.OfflineTtsVitsModelConfig
type OfflineTtsMatchaModelConfig = sherpa.OfflineTtsMatchaModelConfig
type OfflineTtsKokoroModelConfig = sherpa.OfflineTtsKokoroModelConfig
type OfflineTtsKittenModelConfig = sherpa.OfflineTtsKittenModelConfig
type OfflineTtsPocketModelConfig = sherpa.OfflineTtsPocketModelConfig
type OfflineTtsZipvoiceModelConfig = sherpa.OfflineTtsZipvoiceModelConfig
type OfflineTtsModelConfig = sherpa.OfflineTtsModelConfig
type OfflineTtsConfig = sherpa.OfflineTtsConfig
type GeneratedAudio = sherpa.GeneratedAudio
type GenerationConfig = sherpa.GenerationConfig
type OfflineTts = sherpa.OfflineTts
type SileroVadModelConfig = sherpa.SileroVadModelConfig
type TenVadModelConfig = sherpa.TenVadModelConfig
type VadModelConfig = sherpa.VadModelConfig
type CircularBuffer = sherpa.CircularBuffer
type SpeechSegment = sherpa.SpeechSegment
type VoiceActivityDetector = sherpa.VoiceActivityDetector
type SpokenLanguageIdentificationWhisperConfig = sherpa.SpokenLanguageIdentificationWhisperConfig
type SpokenLanguageIdentificationConfig = sherpa.SpokenLanguageIdentificationConfig
type SpokenLanguageIdentification = sherpa.SpokenLanguageIdentification
type SpokenLanguageIdentificationResult = sherpa.SpokenLanguageIdentificationResult
type SpeakerEmbeddingExtractorConfig = sherpa.SpeakerEmbeddingExtractorConfig
type SpeakerEmbeddingExtractor = sherpa.SpeakerEmbeddingExtractor
type SpeakerEmbeddingManager = sherpa.SpeakerEmbeddingManager
type OfflineSpeakerSegmentationPyannoteModelConfig = sherpa.OfflineSpeakerSegmentationPyannoteModelConfig
type OfflineSpeakerSegmentationModelConfig = sherpa.OfflineSpeakerSegmentationModelConfig
type FastClusteringConfig = sherpa.FastClusteringConfig
type OfflineSpeakerDiarizationConfig = sherpa.OfflineSpeakerDiarizationConfig
type OfflineSpeakerDiarization = sherpa.OfflineSpeakerDiarization
type OfflineSpeakerDiarizationSegment = sherpa.OfflineSpeakerDiarizationSegment
type OfflinePunctuationModelConfig = sherpa.OfflinePunctuationModelConfig
type OfflinePunctuationConfig = sherpa.OfflinePunctuationConfig
type OfflinePunctuation = sherpa.OfflinePunctuation
type KeywordSpotterConfig = sherpa.KeywordSpotterConfig
type KeywordSpotterResult = sherpa.KeywordSpotterResult
type KeywordSpotter = sherpa.KeywordSpotter
type OfflineZipformerAudioTaggingModelConfig = sherpa.OfflineZipformerAudioTaggingModelConfig
type AudioTaggingModelConfig = sherpa.AudioTaggingModelConfig
type AudioTaggingConfig = sherpa.AudioTaggingConfig
type AudioTagging = sherpa.AudioTagging
type AudioEvent = sherpa.AudioEvent
type OfflineSpeechDenoiserGtcrnModelConfig = sherpa.OfflineSpeechDenoiserGtcrnModelConfig
type OfflineSpeechDenoiserModelConfig = sherpa.OfflineSpeechDenoiserModelConfig
type OfflineSpeechDenoiserConfig = sherpa.OfflineSpeechDenoiserConfig
type OfflineSpeechDenoiser = sherpa.OfflineSpeechDenoiser
type DenoisedAudio = sherpa.DenoisedAudio
type Wave = sherpa.Wave
var DeleteOnlineRecognizer = sherpa.DeleteOnlineRecognizer
var NewOnlineRecognizer = sherpa.NewOnlineRecognizer
var DeleteOnlineStream = sherpa.DeleteOnlineStream
var NewOnlineStream = sherpa.NewOnlineStream
var DeleteOfflineRecognizer = sherpa.DeleteOfflineRecognizer
var NewOfflineRecognizer = sherpa.NewOfflineRecognizer
var DeleteOfflineStream = sherpa.DeleteOfflineStream
var NewOfflineStream = sherpa.NewOfflineStream
var DeleteOfflineTts = sherpa.DeleteOfflineTts
var NewOfflineTts = sherpa.NewOfflineTts
var DeleteCircularBuffer = sherpa.DeleteCircularBuffer
var NewCircularBuffer = sherpa.NewCircularBuffer
var NewVoiceActivityDetector = sherpa.NewVoiceActivityDetector
var DeleteVoiceActivityDetector = sherpa.DeleteVoiceActivityDetector
var NewSpokenLanguageIdentification = sherpa.NewSpokenLanguageIdentification
var DeleteSpokenLanguageIdentification = sherpa.DeleteSpokenLanguageIdentification
var NewSpeakerEmbeddingExtractor = sherpa.NewSpeakerEmbeddingExtractor
var DeleteSpeakerEmbeddingExtractor = sherpa.DeleteSpeakerEmbeddingExtractor
var NewSpeakerEmbeddingManager = sherpa.NewSpeakerEmbeddingManager
var DeleteSpeakerEmbeddingManager = sherpa.DeleteSpeakerEmbeddingManager
var ReadWave = sherpa.ReadWave
var DeleteOfflineSpeakerDiarization = sherpa.DeleteOfflineSpeakerDiarization
var NewOfflineSpeakerDiarization = sherpa.NewOfflineSpeakerDiarization
var NewOfflinePunctuation = sherpa.NewOfflinePunctuation
var DeleteOfflinePunc = sherpa.DeleteOfflinePunc
var DeleteKeywordSpotter = sherpa.DeleteKeywordSpotter
var NewKeywordSpotter = sherpa.NewKeywordSpotter
var NewKeywordStream = sherpa.NewKeywordStream
var NewKeywordStreamWithKeywords = sherpa.NewKeywordStreamWithKeywords
var DeleteAudioTagging = sherpa.DeleteAudioTagging
var NewAudioTagging = sherpa.NewAudioTagging
var NewAudioTaggingStream = sherpa.NewAudioTaggingStream
var DeleteOfflineSpeechDenoiser = sherpa.DeleteOfflineSpeechDenoiser
var NewOfflineSpeechDenoiser = sherpa.NewOfflineSpeechDenoiser
var GetVersion = sherpa.GetVersion
var GetGitSha1 = sherpa.GetGitSha1
var GetGitDate = sherpa.GetGitDate
