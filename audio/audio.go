package audio

import (
	"io"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AudioResource struct {
	Source    AudioStream
	MusicType MusicType
}

type AudioStream interface {
	io.ReadSeeker
	Length() int64
}

type MusicType int

const (
	TypeOgg MusicType = iota
	TypeMP3
)

func (t MusicType) String() string {
	switch t {
	case TypeOgg:
		return "Ogg"
	case TypeMP3:
		return "MP3"
	default:
		return "unsupported"
	}
}

func LoadAudioFromFile(path string, audioContext *audio.Context, musicType MusicType) (*AudioResource, error) {
	aR := AudioResource{MusicType: musicType}
	var s AudioStream

	fileStream, err := ebitenutil.OpenFile(path)

	if err != nil {
		return nil, err
	}

	switch musicType {
	case TypeOgg:
		var err error
		s, err = vorbis.Decode(audioContext, fileStream)
		if err != nil {
			return nil, err
		}
	case TypeMP3:
		var err error
		s, err = mp3.Decode(audioContext, fileStream)
		if err != nil {
			return nil, err
		}
	default:
		panic("not reached")
	}

	aR.Source = s
	return &aR, nil
}
