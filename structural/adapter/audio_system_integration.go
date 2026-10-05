package adapter

import (
	"fmt"
	"path/filepath"
	"strings"
)

type MediaPlayer interface {
	Play(fileName string)
}

type MP3AudioPlayer struct{}

func (m *MP3AudioPlayer) Play(fileName string) {
	fmt.Printf("Playing MP3 audio file: %s\n", fileName)
}

type WAVAudioPlayer struct{}

func (w *WAVAudioPlayer) Play(fileName string) {
	fmt.Printf("Playing WAV audio file: %s\n", fileName)
}

type LegacyAudioEngine struct{}

func (l *LegacyAudioEngine) PlayAacFile(fileName string) {
	fmt.Printf("Playing AAC file via Legacy Engine: %s\n", fileName)
}

func (l *LegacyAudioEngine) PlayOggFile(fileName string) {
	fmt.Printf("Playing OGG file via Legacy Engine: %s\n", fileName)
}

type LegacyAudioAdapter struct {
	legacyEngine *LegacyAudioEngine
}

func NewLegacyAudioAdapter(engine *LegacyAudioEngine) *LegacyAudioAdapter {
	return &LegacyAudioAdapter{legacyEngine: engine}
}

func (a *LegacyAudioAdapter) Play(fileName string) {
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".aac":
		a.legacyEngine.PlayAacFile(fileName)
	case ".ogg":
		a.legacyEngine.PlayOggFile(fileName)
	default:
		fmt.Printf("Format %s not supported by Legacy Adapter\n", ext)
	}
}

type AudioPlayer struct {
	mp3Player     *MP3AudioPlayer
	wavPlayer     *WAVAudioPlayer
	legacyAdapter *LegacyAudioAdapter
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{
		mp3Player:     &MP3AudioPlayer{},
		wavPlayer:     &WAVAudioPlayer{},
		legacyAdapter: NewLegacyAudioAdapter(&LegacyAudioEngine{}),
	}
}

func (a *AudioPlayer) Play(fileName string) {
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".mp3":
		a.mp3Player.Play(fileName)
	case ".wav":
		a.wavPlayer.Play(fileName)
	case ".aac", ".ogg":
		a.legacyAdapter.Play(fileName)
	default:
		fmt.Printf("Unsupported media format: %s\n", ext)
	}
}
