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
	Repo       repository.Repository
	SpotifyAPI repository.Spotify
}

func NewMusicService(repo repository.Repository, spotify repository.Spotify) *music {
	return &music{Repo: repo, SpotifyAPI: spotify}
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

	/*
		TODO
		HERE THE MODEL ENDPOINT GOES
	*/

	track.Mood = 4

	userTrack, err := u.Repo.StoreTrack(ctx, track)
	if err != nil {
		return nil, err
	}
	return userTrack, nil
}
