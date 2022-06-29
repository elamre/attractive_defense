package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

//go:embed sounds/insufficient_funds.ogg
var insufficientFunds []byte

//go:embed sounds/research_complete.ogg
var researchComplete []byte

//go:embed sounds/new_wave.ogg
var newWave []byte

//go:embed sounds/your_base.ogg
var yourBase []byte

//go:embed sounds/Jammin.mp3
var background []byte

type SoundManager struct {
	backgroundMusic             *audio.Player
	insufficientFundsPlayer     *audio.Player
	researchCompleteFundsPlayer *audio.Player
	waveCompleteFundsPlayer     *audio.Player
	attackCompleteFundsPlayer   *audio.Player
	context                     *audio.Context

	backgroundStream *mp3.Stream
	backgroundPlayer *audio.Player

	//backgroundContext *audio.Context
}

var StaticSoundManager *SoundManager

func NewSoundManager() *SoundManager {
	var err error
	ss := SoundManager{}
	ss.context = audio.NewContext(44100)
	//ss.backgroundContext = audio.NewContext(44100)
	insufficientFundsStream, err := vorbis.DecodeWithSampleRate(44100, bytes.NewReader(insufficientFunds))
	if err != nil {
		panic(err)
	}
	player, err := ss.context.NewPlayer(insufficientFundsStream)
	if err != nil {
		panic(err)
	}
	ss.insufficientFundsPlayer = player

	researchCompleteStream, err := vorbis.DecodeWithSampleRate(44100, bytes.NewReader(researchComplete))
	if err != nil {
		panic(err)
	}
	player, err = ss.context.NewPlayer(researchCompleteStream)
	if err != nil {
		panic(err)
	}
	ss.researchCompleteFundsPlayer = player

	newWaveStream, err := vorbis.DecodeWithSampleRate(44100, bytes.NewReader(newWave))
	if err != nil {
		panic(err)
	}
	player, err = ss.context.NewPlayer(newWaveStream)
	if err != nil {
		panic(err)
	}
	ss.waveCompleteFundsPlayer = player

	yourBaseStream, err := vorbis.DecodeWithSampleRate(44100, bytes.NewReader(yourBase))
	if err != nil {
		panic(err)
	}
	player, err = ss.context.NewPlayer(yourBaseStream)
	if err != nil {
		panic(err)
	}
	ss.attackCompleteFundsPlayer = player

	ss.backgroundStream, err = mp3.DecodeWithSampleRate(44100, bytes.NewReader(background))
	if err != nil {
		panic(err)
	}
	player, err = ss.context.NewPlayer(ss.backgroundStream)
	if err != nil {
		panic(err)
	}
	ss.backgroundPlayer = player
	ss.backgroundPlayer.SetVolume(0.1)
	ss.backgroundPlayer.Play()
	return &ss
}

func (s *SoundManager) Update() {
	if s.backgroundPlayer.Current() >= 184372244000 {
		err := s.backgroundPlayer.Rewind()
		if err != nil {
			panic(err)
		}
		s.backgroundPlayer.Play()
	}
}

func (s *SoundManager) PlayInsufficientFunds() {
	if s.insufficientFundsPlayer.Current() >= 1509297052 {
		_ = s.insufficientFundsPlayer.Rewind()
	}
	if !s.insufficientFundsPlayer.IsPlaying() {
		s.insufficientFundsPlayer.Play()
	}
}

func (s *SoundManager) PlayNewWave() {
	if s.waveCompleteFundsPlayer.Current() >= 1021678004 {
		_ = s.waveCompleteFundsPlayer.Rewind()
	}
	if !s.waveCompleteFundsPlayer.IsPlaying() {
		s.waveCompleteFundsPlayer.Play()
	}
}

func (s *SoundManager) PlayYourBaseUnderAttack() {
	if s.attackCompleteFundsPlayer.Current() >= 1532517006 {
		_ = s.attackCompleteFundsPlayer.Rewind()
	}
	if !s.attackCompleteFundsPlayer.IsPlaying() {
		s.attackCompleteFundsPlayer.Play()
	}
}
func (s *SoundManager) PlayResearchComplete() {
	if s.researchCompleteFundsPlayer.Current() >= 940408163 {
		_ = s.researchCompleteFundsPlayer.Rewind()
	}
	if !s.researchCompleteFundsPlayer.IsPlaying() {
		s.researchCompleteFundsPlayer.Play()
	}
}
