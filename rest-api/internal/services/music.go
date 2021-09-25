package services

import (
	"context"
	"fmt"
	"rest-api/internal/models"
	"rest-api/internal/repository"
)

type Music interface {
	// Store stores the input
	Store(ctx context.Context, user *models.Track) (*models.Track, error)
}

type music struct {
	Repo        repository.Repository
	SpotifyAPI  repository.Spotify
	TrainingAPI repository.Training
}

func NewMusicService(repo repository.Repository, spotify repository.Spotify, training repository.Training) *music {
	return &music{Repo: repo, SpotifyAPI: spotify, TrainingAPI: training}
}

func (u music) Store(ctx context.Context, track *models.Track) (*models.Track, error) {
	audioFeatures, err := u.SpotifyAPI.GetAudioAnalysisByTrackID(ctx, track.TrackID)
	if err != nil {
		return nil, err
	}

	track.Features = fmt.Sprintf("%f,%f,%f,%f,%f,%f,%f,%f,%f,%f",
		float32(audioFeatures.Duration),
		audioFeatures.Danceability,
		audioFeatures.Acousticness,
		audioFeatures.Energy,
		audioFeatures.Instrumentalness,
		audioFeatures.Liveness,
		audioFeatures.Valence,
		audioFeatures.Loudness,
		audioFeatures.Speechiness,
		audioFeatures.Tempo)

	mood, err := u.TrainingAPI.GetMood(ctx, track.Features)
	if err != nil {
		return nil, err
	}
	track.Mood = mood

	userTrack, err := u.Repo.StoreTrack(ctx, track)
	if err != nil {
		return nil, err
	}
	return userTrack, nil
}
