package services

import (
	"anime/internal/models"

	"github.com/google/uuid"
)

type Anime struct {
	ID           string      `json:"id,omitempty"`
	Language     string      `json:"lang"`
	Title        string      `json:"title"`
	Genre        string      `json:"genre"`
	EpisodeCount int         `json:"episodes"`
	DateRelease  string      `json:"release"`
	Subtitles    []Subtitles `json:"sub"`
}

type Subtitles struct {
	ID       string `json:"id,omitempty"`
	Language string `json:"lang"`
	Author   string `json:"autor"`
}

type AnimeService interface {
	CreateTitle()
	UpdateTitle()
	DeleteTitle()
	Title()
	Titles()
}

type service struct {
	storage string
	logger  string
	config  string
}

func NewService(store, log, cfg string) AnimeService {
	return &service{
		storage: store,
		logger:  log,
		config:  cfg,
	}
}

func (s *service) CreateTitle() {

}

func (s *service) UpdateTitle() {

}

func (s *service) DeleteTitle() {

}

func (s *service) Title() {

}

func (s *service) Titles() {

}

// todo: написать реализпцию функции
func toModelData(anime Anime) models.Anime {
	id := uuid.New()

	sub := make([]models.Subtitles, len(anime.Subtitles))
	// for _, value := range anime.Subtitles {
	// 	var convert models.Subtitles
	// 	id := uuid.New()
	// 	convert = models.Subtitles{
	// 		ID:         id,
	// 		Created_at: time.Now(),
	// 		Enable:     true,
	// 		Language:   value.Language,
	// 		Author:     value.Author,
	// 	}
	// 	sub = append(sub, convert)
	// }
	// todo:написать цикл конвертации

	for i, v := range anime.Subtitles {
		subID := uuid.New()
		sub[i] = models.Subtitles{
			ID:       subID,
			Language: v.Language,
			Author:   v.Author,
			Enable:   true,
		}
	}

	return models.Anime{
		ID:        id,
		Language:  anime.Language,
		Subtitles: sub,
	}
}
